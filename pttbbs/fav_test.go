// Copyright 2020 Pichu Chen, The PTT APP Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pttbbs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	if line.LineID != 2 {
		t.Errorf("Line should have ID 2 but was %d", line.LineID)
	}

	encoded, err := item.MarshalBinary()
	if bytes.Compare(data, encoded) != 0 {
		t.Errorf("Encoded FavLineItem should be equal to the input")
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
					LineID:   0,
					FolderID: 0,
					FavItems: []*FavItem{
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardID:   14,
								LastVisit: time0,
								Attr:      0,
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardID:   6,
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
					LineID:   0,
					FolderID: 1,
					FavItems: []*FavItem{
						{
							FavType: 2,
							FavAttr: 1,
							Item: &FavFolderItem{
								FolderID: 1,
								Title:    "新的目錄",
								ThisFolder: &FavFolder{
									NAlloc:   11,
									DataTail: 3,
									NBoards:  0,
									NLines:   2,
									NFolders: 1,
									LineID:   2,
									FolderID: 1,
									FavItems: []*FavItem{
										{
											FavType: 3,
											FavAttr: 1,
											Item:    &FavLineItem{LineID: 1},
										},
										{
											FavType: 2,
											FavAttr: 1,
											Item: &FavFolderItem{
												FolderID: 1,
												Title:    "Folder02",
												ThisFolder: &FavFolder{
													NAlloc:   9,
													DataTail: 1,
													NBoards:  0,
													NLines:   0,
													NFolders: 1,
													LineID:   0,
													FolderID: 1,
													FavItems: []*FavItem{
														{
															FavType: 2,
															FavAttr: 1,
															Item: &FavFolderItem{
																FolderID: 1,
																Title:    "Folder03",
																ThisFolder: &FavFolder{
																	NAlloc:   12,
																	DataTail: 4,
																	NBoards:  1,
																	NLines:   1,
																	NFolders: 2,
																	LineID:   1,
																	FolderID: 2,
																	FavItems: []*FavItem{
																		{
																			FavType: 2,
																			FavAttr: 1,
																			Item: &FavFolderItem{
																				FolderID: 1,
																				Title:    "Folder04",
																				ThisFolder: &FavFolder{
																					NAlloc:   9,
																					DataTail: 1,
																					NBoards:  1,
																					NLines:   0,
																					NFolders: 0,
																					LineID:   0,
																					FolderID: 0,
																					FavItems: []*FavItem{
																						{
																							FavType: 1,
																							FavAttr: 1,
																							Item: &FavBoardItem{
																								BoardID:   14,
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
																				BoardID:   14,
																				LastVisit: time0,
																				Attr:      0,
																			},
																		},
																		{
																			FavType: 3,
																			FavAttr: 1,
																			Item:    &FavLineItem{LineID: 1},
																		},
																		{
																			FavType: 2,
																			FavAttr: 1,
																			Item: &FavFolderItem{
																				FolderID: 2,
																				Title:    "MAX Length:2345672234567890323456789042345678905",
																				ThisFolder: &FavFolder{
																					NAlloc:   8,
																					DataTail: 0,
																					NBoards:  0,
																					NLines:   0,
																					NFolders: 0,
																					LineID:   0,
																					FolderID: 0,
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
											Item:    &FavLineItem{LineID: 2},
										},
									},
								},
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardID:   13,
								LastVisit: time0,
								Attr:      0,
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardID:   14,
								LastVisit: time0,
								Attr:      0,
							},
						},
						{
							FavType: 1,
							FavAttr: 1,
							Item: &FavBoardItem{
								BoardID:   6,
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
		filePath := fmt.Sprintf("testcase/fav/%s", c.filename)
		fav, err := OpenFavFile(filePath)
		if err != nil {
			t.Fatal(err)
		}
		favJSON, _ := json.Marshal(fav)
		expectedJSON, _ := json.Marshal(c.expected)
		if strings.Compare(string(favJSON), string(expectedJSON)) != 0 {
			t.Errorf("Parsed result does not equal to expected. Expected json string: %s \nParsed json string: %s", string(expectedJSON), string(favJSON))
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

		expectedBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			t.Fatal(err)
		}
		encoded, err := fav.MarshalBinary()
		if bytes.Compare(expectedBytes, encoded) != 0 {
			t.Errorf("Encoded FavFile should be equal to the input. Expected: \n%v\nEncoded: \n%v\n ", expectedBytes, encoded)
		}
	}

}
