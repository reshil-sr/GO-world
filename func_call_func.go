package main

import "fmt"

func CallMe() {
	print(area, 2, 4)
	print(sum, 2, 4)
}

func print(f func(int, int) int, a, b int) {
	res := f(a, b)
	fmt.Println(res)
}

func area(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}

//=========================
func (c *S3) ListObjectsV2PagesWithContext(ctx aws.Context, input *ListObjectsV2Input, fn func(*ListObjectsV2Output, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListObjectsV2Input
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListObjectsV2Request(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListObjectsV2Output), !p.HasNextPage())
	}
	return p.Err()
}
mockS3 := mockS3API{
	MockListObjectsV2PagesWithContext: func(ctx aws.Context, input *s3.ListObjectsV2Input, fn func(output *s3.ListObjectsV2Output, lastPage bool) bool, opts ...request.Option) error {
		op := &s3.ListObjectsV2Output{
			Contents: []*s3.Object{
				{
					Key:          aws.String("abcd.jpg"),
					ETag:         aws.String("abcdefg"),
					Size:         aws.Int64(10),
					LastModified: aws.Time(lastModifiedTime),
				},
			},
			CommonPrefixes: []*s3.CommonPrefix{
				{
					Prefix: aws.String("a/b/c"),
				},
			},
			IsTruncated: aws.Bool(false),
		}
		fn(op, true)
		return nil
	},
}


type mockS3V2API struct {
	MockNewListObjectsV2Paginator func(client *awsS3V2.Client, params *awsS3V2.ListObjectsV2Input, optFns ...func(*awsS3V2.ListObjectsV2PaginatorOptions)) *awsS3V2.ListObjectsV2Paginator
}

// ListObjectsV2PagesWithContext is used to call the mocked version of the same in tests
func (m mockS3V2API) NewListObjectsV2Paginator(client *awsS3V2.Client, params *awsS3V2.ListObjectsV2Input, optFns ...func(*awsS3V2.ListObjectsV2PaginatorOptions)) *awsS3V2.ListObjectsV2Paginator {
	return m.MockNewListObjectsV2Paginator(client, params, optFns...)
}

type mockListObjectsV2Pager struct {
	PageNum int
	Pages   []*s3.ListObjectsV2Output
}

func (m *mockListObjectsV2Pager) HasMorePages() bool {
	return m.PageNum < len(m.Pages)
}

func (m *mockListObjectsV2Pager) NextPage(ctx context.Context, f ...func(*s3.Options)) (output *s3.ListObjectsV2Output, err error) {
	if m.PageNum >= len(m.Pages) {
		return nil, fmt.Errorf("no more pages")
	}
	output = m.Pages[m.PageNum]
	m.PageNum++
	return output, nil
}

func TestCountObjects(t *testing.T) {
	pager := &mockListObjectsV2Pager{
		Pages: []*s3.ListObjectsV2Output{
			{
				Contents: []*s3.Object{
					{
						Key:          aws.String("abcd.jpg"),
						ETag:         aws.String("abcdefg"),
						Size:         aws.Int64(10),
						LastModified: aws.Time(lastModifiedTime),
					},
				},
				IsTruncated: aws.Bool(false),
			},
		},
	}
	objects, err := CountObjects(context.TODO(), pager)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}
	if expect, actual := 30, objects; expect != actual {
		t.Errorf("expect %v, got %v", expect, actual)
	}
}
================================================================
type TestIface interface {
	NewListObjectsV2Paginator(client *awsS3V2.ListObjectsV2APIClient, params *awsS3V2.ListObjectsV2Input, optFns ...func(*awsS3V2.ListObjectsV2PaginatorOptions)) *awsS3V2.ListObjectsV2Paginator
}
type TestStruct struct {
	MockNewListObjectsV2Paginator func(client *awsS3V2.ListObjectsV2APIClient, params *awsS3V2.ListObjectsV2Input, optFns ...func(*awsS3V2.ListObjectsV2PaginatorOptions)) *awsS3V2.ListObjectsV2Paginator
}

func (m *TestStruct) NewListObjectsV2Paginator(client *awsS3V2.ListObjectsV2APIClient, params *awsS3V2.ListObjectsV2Input, optFns ...func(*awsS3V2.ListObjectsV2PaginatorOptions)) *awsS3V2.ListObjectsV2Paginator {
	return m.MockNewListObjectsV2Paginator(client, params, optFns...)
}