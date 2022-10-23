package plan

import (
	"log"
	"testing"

	"github.com/wangtao334/ak47/controller"
	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/group"
	"github.com/wangtao334/ak47/sampler"
	"github.com/wangtao334/ak47/variable"
)

func TestTestPlan_Run(t1 *testing.T) {
	t := &TestPlan{
		Name: "my test plan",
		Children: []element.Element{
			&group.Group{
				Parent: &element.Parent{
					Name:   "group1",
					Enable: true,
					Children: []element.Element{
						&variable.Text{
							Parent: &element.Parent{
								Name:   "age1",
								Enable: true,
							},
							Value: "17",
						},
						&variable.CSV{
							Parent: &element.Parent{
								Enable: true,
							},
							FilePath: `F:/csv2`,
						},
						&controller.Once{
							Parent: &element.Parent{
								Enable: true,
								Children: []element.Element{
									&variable.RandomText{
										Parent:&element.Parent{
											Name: "rand text",
											Enable: true,
										},
										Len: "10",
										Set: "中国人",
									},
								},
							},
						},
						&controller.Loops{
							Parent: &element.Parent{
								Name:   "loops controller1",
								Enable: true,
								Children: []element.Element{
									&variable.Text{
										Parent: &element.Parent{
											Name:   "name1",
											Enable: true,
										},
										Value: "su cui",
									},
									&sampler.Demo{
										Parent:&element.Parent{
											Name: "demo1",
											Enable: true,
										},
										Url: &variable.Text{
											Parent:&element.Parent{
												Enable: true,
											},
											Value: "http://${host}/test?name=${name1}&age=${age1}",
										},
										Headers: []element.Element{
											&variable.Text{
												Parent: &element.Parent{
													Name: "lang",
													Enable: true,
												},
												Value: "en",
											},
											&variable.Text{
												Parent: &element.Parent{
													Name: "lang",
													Enable: true,
												},
												Value: "cn",
											},
										},
										//Body: &variable.File{
										//	Parent:&element.Parent{
										//		Enable: true,
										//	},
										//	FilePath: `F:/file1`,
										//},
										//Body: &variable.RandomText{
										//	Parent:&element.Parent{
										//		Enable: true,
										//	},
										//	Len: "10",
										//	Set: "中国人",
										//},
										Body: &variable.Text{
											Parent:&element.Parent{
												Enable: true,
											},
											Value: "${rand text}",
										},
									},
								},
							},
							Loops: "2",
						},
					},
				},
				Mode:      1,
				WorkerNum: "1",
				Loops:     "${loops}",
			},
			&variable.Text{
				Parent: &element.Parent{
					Name:   "name",
					Enable: true,
				},
				Value: "wang tao",
			},
			&variable.Text{
				Parent: &element.Parent{
					Name:   "host",
					Enable: true,
				},
				Value: "localhost",
			},
			&variable.Text{
				Parent: &element.Parent{
					Name:   "age",
					Enable: true,
				},
				Value: "18",
			},
			&variable.Text{
				Parent: &element.Parent{
					Name:   "desc",
					Enable: true,
				},
				Value: "${name}-${age}",
			},
			&variable.Text{
				Parent: &element.Parent{
					Name:   "loops",
					Enable: true,
				},
				Value: "3",
			},
			&variable.CSV{
				Parent: &element.Parent{
					Enable: true,
				},
				FilePath: `F:/csv1`,
			},
			&variable.File{
				Parent: &element.Parent{
					Name: "f",
					Enable: true,
				},
				FilePath: `F:/file2`,
			},
		},
	}
	if err := t.Run(); err != nil {
		log.Println(err)
	}
}
