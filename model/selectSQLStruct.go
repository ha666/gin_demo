package model

type SelectSQLStruct struct {
	Cache       string      `json:"Cache"`
	Comments    interface{} `json:"Comments"`
	Distinct    string      `json:"Distinct"`
	Hints       string      `json:"Hints"`
	SelectExprs []struct {
		Expr struct {
			Metadata interface{} `json:"Metadata"`
			Name     string      `json:"Name"`
		} `json:"Expr"`
		As string `json:"As"`
	} `json:"SelectExprs"`
	From []struct {
		LeftExpr struct {
			Expr struct {
				Name      string `json:"Name"`
				Qualifier string `json:"Qualifier"`
			} `json:"Expr"`
			Partitions interface{} `json:"Partitions"`
			As         string      `json:"As"`
			Hints      interface{} `json:"Hints"`
			LeftExpr   struct {
				Expr struct {
					Name      string `json:"Name"`
					Qualifier string `json:"Qualifier"`
				} `json:"Expr"`
				Partitions interface{} `json:"Partitions"`
				As         string      `json:"As"`
				Hints      interface{} `json:"Hints"`
				LeftExpr   struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
									LeftExpr   struct {
										Expr struct {
											Name      string `json:"Name"`
											Qualifier string `json:"Qualifier"`
										} `json:"Expr"`
										Partitions interface{} `json:"Partitions"`
										As         string      `json:"As"`
										Hints      interface{} `json:"Hints"`
										LeftExpr   struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
											LeftExpr   struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"LeftExpr"`
											Join      string `json:"Join"`
											RightExpr struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"RightExpr"`
										} `json:"LeftExpr"`
										Join      string `json:"Join"`
										RightExpr struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
											LeftExpr   struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"LeftExpr"`
											Join      string `json:"Join"`
											RightExpr struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"RightExpr"`
										} `json:"RightExpr"`
									} `json:"LeftExpr"`
									Join      string `json:"Join"`
									RightExpr struct {
										Expr struct {
											Name      string `json:"Name"`
											Qualifier string `json:"Qualifier"`
										} `json:"Expr"`
										Partitions interface{} `json:"Partitions"`
										As         string      `json:"As"`
										Hints      interface{} `json:"Hints"`
										LeftExpr   struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
											LeftExpr   struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"LeftExpr"`
											Join      string `json:"Join"`
											RightExpr struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"RightExpr"`
										} `json:"LeftExpr"`
										Join      string `json:"Join"`
										RightExpr struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
											LeftExpr   struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"LeftExpr"`
											Join      string `json:"Join"`
											RightExpr struct {
												Expr struct {
													Name      string `json:"Name"`
													Qualifier string `json:"Qualifier"`
												} `json:"Expr"`
												Partitions interface{} `json:"Partitions"`
												As         string      `json:"As"`
												Hints      interface{} `json:"Hints"`
											} `json:"RightExpr"`
										} `json:"RightExpr"`
									} `json:"RightExpr"`
								} `json:"RightExpr"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
									LeftExpr   struct {
										Expr struct {
											Name      string `json:"Name"`
											Qualifier string `json:"Qualifier"`
										} `json:"Expr"`
										Partitions interface{} `json:"Partitions"`
										As         string      `json:"As"`
										Hints      interface{} `json:"Hints"`
										LeftExpr   struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"LeftExpr"`
										Join      string `json:"Join"`
										RightExpr struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"RightExpr"`
									} `json:"LeftExpr"`
									Join      string `json:"Join"`
									RightExpr struct {
										Expr struct {
											Name      string `json:"Name"`
											Qualifier string `json:"Qualifier"`
										} `json:"Expr"`
										Partitions interface{} `json:"Partitions"`
										As         string      `json:"As"`
										Hints      interface{} `json:"Hints"`
										LeftExpr   struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"LeftExpr"`
										Join      string `json:"Join"`
										RightExpr struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"RightExpr"`
									} `json:"RightExpr"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
									LeftExpr   struct {
										Expr struct {
											Name      string `json:"Name"`
											Qualifier string `json:"Qualifier"`
										} `json:"Expr"`
										Partitions interface{} `json:"Partitions"`
										As         string      `json:"As"`
										Hints      interface{} `json:"Hints"`
										LeftExpr   struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"LeftExpr"`
										Join      string `json:"Join"`
										RightExpr struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"RightExpr"`
									} `json:"LeftExpr"`
									Join      string `json:"Join"`
									RightExpr struct {
										Expr struct {
											Name      string `json:"Name"`
											Qualifier string `json:"Qualifier"`
										} `json:"Expr"`
										Partitions interface{} `json:"Partitions"`
										As         string      `json:"As"`
										Hints      interface{} `json:"Hints"`
										LeftExpr   struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"LeftExpr"`
										Join      string `json:"Join"`
										RightExpr struct {
											Expr struct {
												Name      string `json:"Name"`
												Qualifier string `json:"Qualifier"`
											} `json:"Expr"`
											Partitions interface{} `json:"Partitions"`
											As         string      `json:"As"`
											Hints      interface{} `json:"Hints"`
										} `json:"RightExpr"`
									} `json:"RightExpr"`
								} `json:"RightExpr"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"RightExpr"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"RightExpr"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"RightExpr"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"RightExpr"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"RightExpr"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
								LeftExpr   struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"LeftExpr"`
								Join      string `json:"Join"`
								RightExpr struct {
									Expr struct {
										Name      string `json:"Name"`
										Qualifier string `json:"Qualifier"`
									} `json:"Expr"`
									Partitions interface{} `json:"Partitions"`
									As         string      `json:"As"`
									Hints      interface{} `json:"Hints"`
								} `json:"RightExpr"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"LeftExpr"`
				Join      string `json:"Join"`
				RightExpr struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"RightExpr"`
			} `json:"LeftExpr"`
			Join      string `json:"Join"`
			RightExpr struct {
				Expr struct {
					Name      string `json:"Name"`
					Qualifier string `json:"Qualifier"`
				} `json:"Expr"`
				Partitions interface{} `json:"Partitions"`
				As         string      `json:"As"`
				Hints      interface{} `json:"Hints"`
				LeftExpr   struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"LeftExpr"`
				Join      string `json:"Join"`
				RightExpr struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"RightExpr"`
			} `json:"RightExpr"`
		} `json:"LeftExpr"`
		Join      string `json:"Join"`
		RightExpr struct {
			Expr struct {
				Name      string `json:"Name"`
				Qualifier string `json:"Qualifier"`
			} `json:"Expr"`
			Partitions interface{} `json:"Partitions"`
			As         string      `json:"As"`
			Hints      interface{} `json:"Hints"`
			LeftExpr   struct {
				Expr struct {
					Name      string `json:"Name"`
					Qualifier string `json:"Qualifier"`
				} `json:"Expr"`
				Partitions interface{} `json:"Partitions"`
				As         string      `json:"As"`
				Hints      interface{} `json:"Hints"`
				LeftExpr   struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"LeftExpr"`
				Join      string `json:"Join"`
				RightExpr struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"RightExpr"`
			} `json:"LeftExpr"`
			Join      string `json:"Join"`
			RightExpr struct {
				Expr struct {
					Name      string `json:"Name"`
					Qualifier string `json:"Qualifier"`
				} `json:"Expr"`
				Partitions interface{} `json:"Partitions"`
				As         string      `json:"As"`
				Hints      interface{} `json:"Hints"`
				LeftExpr   struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"LeftExpr"`
				Join      string `json:"Join"`
				RightExpr struct {
					Expr struct {
						Name      string `json:"Name"`
						Qualifier string `json:"Qualifier"`
					} `json:"Expr"`
					Partitions interface{} `json:"Partitions"`
					As         string      `json:"As"`
					Hints      interface{} `json:"Hints"`
					LeftExpr   struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"LeftExpr"`
					Join      string `json:"Join"`
					RightExpr struct {
						Expr struct {
							Name      string `json:"Name"`
							Qualifier string `json:"Qualifier"`
						} `json:"Expr"`
						Partitions interface{} `json:"Partitions"`
						As         string      `json:"As"`
						Hints      interface{} `json:"Hints"`
						LeftExpr   struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"LeftExpr"`
						Join      string `json:"Join"`
						RightExpr struct {
							Expr struct {
								Name      string `json:"Name"`
								Qualifier string `json:"Qualifier"`
							} `json:"Expr"`
							Partitions interface{} `json:"Partitions"`
							As         string      `json:"As"`
							Hints      interface{} `json:"Hints"`
							LeftExpr   struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"LeftExpr"`
							Join      string `json:"Join"`
							RightExpr struct {
								Expr struct {
									Name      string `json:"Name"`
									Qualifier string `json:"Qualifier"`
								} `json:"Expr"`
								Partitions interface{} `json:"Partitions"`
								As         string      `json:"As"`
								Hints      interface{} `json:"Hints"`
							} `json:"RightExpr"`
						} `json:"RightExpr"`
					} `json:"RightExpr"`
				} `json:"RightExpr"`
			} `json:"RightExpr"`
		} `json:"RightExpr"`
		Expr struct {
			Name      string `json:"Name"`
			Qualifier string `json:"Qualifier"`
		} `json:"Expr"`
		Partitions interface{} `json:"Partitions"`
		As         string      `json:"As"`
		Hints      interface{} `json:"Hints"`
	} `json:"From"`
	GroupBy interface{} `json:"GroupBy"`
	Having  interface{} `json:"Having"`
	OrderBy interface{} `json:"OrderBy"`
	Limit   struct {
		Offset struct {
			Type int    `json:"Type"`
			Val  string `json:"Val"`
		} `json:"Offset"`
		Rowcount struct {
			Type int    `json:"Type"`
			Val  string `json:"Val"`
		} `json:"Rowcount"`
	} `json:"Limit"`
	Lock string `json:"Lock"`
}
