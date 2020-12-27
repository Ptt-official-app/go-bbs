package bbs

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

var time0 = time.Unix(int64(0), 0)

func TestNewFavLineItem(t *testing.T) {
	data := []byte{3, 1, 2}
	item, _, err := NewFavItem(data, 0)
	if err != nil || item == nil {
		t.Fatal("Failed to parse Line")
	}
	line, ok := item.Item.(*FavLineItem)
	if !ok {
		t.Error("Parsed item should be a FavLineItem")
	}
	if item.FavType != FavItemTypeLine {
		t.Errorf("Line type should be FavItemTypeLine but was %d", item.FavType)
	}
	if line.LineId != 2 {
		t.Errorf("Line should have ID 2 but was %d", line.LineId)
	}
}

func TestOpenFavFile(t *testing.T) {
	type testCase struct {
		filename string
		expected *FavFile
	}
	testCases := []*testCase{
		{
			filename: "01.fav",
			expected: &FavFile{
				Version: 3363,
				Folder: &FavFolder{
					NAlloc:   10,
					DataTail: 2,
					NBoards:  2,
					NLines:   0,
					NFolders: 0,
					LineId:   0,
					FolderId: 0,
					FavItems: []*FavItem{
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardId:   14,
								LastVisit: time0,
								Attr:      0,
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardId:   6,
								LastVisit: time0,
								Attr:      0,
							},
						},
					},
				},
			},
		},
		{
			filename: "02.fav",
			expected: &FavFile{
				Version: 3363,
				Folder: &FavFolder{
					NAlloc:   12,
					DataTail: 4,
					NBoards:  3,
					NLines:   0,
					NFolders: 1,
					LineId:   0,
					FolderId: 1,
					FavItems: []*FavItem{
						{
							FavType: 2,
							FavAttr: 1,
							Item: &FavFolderItem{
								FolderId: 1,
								Title:    "新的目錄",
								ThisFolder: &FavFolder{
									NAlloc:   11,
									DataTail: 3,
									NBoards:  0,
									NLines:   2,
									NFolders: 1,
									LineId:   2,
									FolderId: 1,
									FavItems: []*FavItem{
										{
											FavType: 3,
											FavAttr: 1,
											Item:    &FavLineItem{LineId: 1},
										},
										{
											FavType: 2,
											FavAttr: 1,
											Item: &FavFolderItem{
												FolderId: 1,
												Title:    "Folder02",
												ThisFolder: &FavFolder{
													NAlloc:   9,
													DataTail: 1,
													NBoards:  0,
													NLines:   0,
													NFolders: 1,
													LineId:   0,
													FolderId: 1,
													FavItems: []*FavItem{
														{
															FavType: 2,
															FavAttr: 1,
															Item: &FavFolderItem{
																FolderId: 1,
																Title:    "Folder03",
																ThisFolder: &FavFolder{
																	NAlloc:   12,
																	DataTail: 4,
																	NBoards:  1,
																	NLines:   1,
																	NFolders: 2,
																	LineId:   1,
																	FolderId: 2,
																	FavItems: []*FavItem{
																		{
																			FavType: 2,
																			FavAttr: 1,
																			Item: &FavFolderItem{
																				FolderId: 1,
																				Title:    "Folder04",
																				ThisFolder: &FavFolder{
																					NAlloc:   9,
																					DataTail: 1,
																					NBoards:  1,
																					NLines:   0,
																					NFolders: 0,
																					LineId:   0,
																					FolderId: 0,
																					FavItems: []*FavItem{
																						{
																							FavType: 1,
																							FavAttr: 1,
																							Item: &FavBoardItem{
																								BoardId:   14,
																								LastVisit: time0,
																								Attr:      0,
																							},
																						},
																					},
																				},
																			},
																		},
																		{
																			FavType: 1,
																			FavAttr: 1,
																			Item: &FavBoardItem{
																				BoardId:   14,
																				LastVisit: time0,
																				Attr:      0,
																			},
																		},
																		{
																			FavType: 3,
																			FavAttr: 1,
																			Item:    &FavLineItem{LineId: 1},
																		},
																		{
																			FavType: 2,
																			FavAttr: 1,
																			Item: &FavFolderItem{
																				FolderId: 2,
																				Title:    "MAX Length:2345672234567890323456789042345678905",
																				ThisFolder: &FavFolder{
																					NAlloc:   8,
																					DataTail: 0,
																					NBoards:  0,
																					NLines:   0,
																					NFolders: 0,
																					LineId:   0,
																					FolderId: 0,
																					FavItems: []*FavItem{},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
										{
											FavType: 3,
											FavAttr: 1,
											Item:    &FavLineItem{LineId: 2},
										},
									},
								},
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardId:   13,
								LastVisit: time0,
								Attr:      0,
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardId:   14,
								LastVisit: time0,
								Attr:      0,
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardId:   6,
								LastVisit: time0,
								Attr:      0,
							},
						},
					},
				},
			},
		},
	}

	for _, c := range testCases {
		fav, err := OpenFavFile(fmt.Sprintf("testcase/fav/%s", c.filename))
		if err != nil {
			t.Fatal(err)
		}
		favJson, _ := json.Marshal(fav)
		expectedJson, _ := json.Marshal(c.expected)
		if strings.Compare(string(favJson), string(expectedJson)) != 0 {
			t.Errorf("Parsed result does not equal to expected. Expected json string: %s \nParsed json string: %s", string(expectedJson), string(favJson))
		}
		for _, item := range fav.Folder.FavItems {
			switch item.FavType {
			case FavItemTypeBoard:
				if target := item.GetBoard(); target == nil {
					t.Error("Fav item type is Board should be able to GetBoard()")
				}
			case FavItemTypeFolder:
				if target := item.GetFolder(); target == nil {
					t.Error("Fav item type is Folder should be able to GetFolder()")
				}
			case FavItemTypeLine:
				if target := item.GetLine(); target == nil {
					t.Error("Fav item type is Line should be able to GetLine()")
				}
			default:
				t.Error("invalid fav type")
			}
		}
	}

}
