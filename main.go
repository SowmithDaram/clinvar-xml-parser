package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type ClinVarDataRelease struct {
	XMLName                   xml.Name           `xml:"ClinVarVariationRelease"`
	Text                      string             `xml:",chardata"`
	Xsi                       string             `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string             `xml:"noNamespaceSchemaLocation,attr"`
	ReleaseDate               string             `xml:"ReleaseDate,attr"`
	Variants                  []VariationArchive `xml:"VariationArchive"`
}

type VariationArchive struct {
	Text                string `xml:",chardata"`
	VariationID         string `xml:"VariationID,attr"`
	VariationName       string `xml:"VariationName,attr"`
	VariationType       string `xml:"VariationType,attr"`
	DateCreated         string `xml:"DateCreated,attr"`
	DateLastUpdated     string `xml:"DateLastUpdated,attr"`
	Accession           string `xml:"Accession,attr"`
	Version             string `xml:"Version,attr"`
	RecordType          string `xml:"RecordType,attr"`
	NumberOfSubmissions string `xml:"NumberOfSubmissions,attr"`
	NumberOfSubmitters  string `xml:"NumberOfSubmitters,attr"`
	RecordStatus        string `xml:"RecordStatus"`
	Species             string `xml:"Species"`
	InterpretedRecord   struct {
		Text         string `xml:",chardata"`
		SimpleAllele struct {
			Text        string `xml:",chardata"`
			AlleleID    string `xml:"AlleleID,attr"`
			VariationID string `xml:"VariationID,attr"`
			GeneList    struct {
				Text string `xml:",chardata"`
				Gene []struct {
					Text             string `xml:",chardata"`
					Symbol           string `xml:"Symbol,attr"`
					FullName         string `xml:"FullName,attr"`
					GeneID           string `xml:"GeneID,attr"`
					HGNCID           string `xml:"HGNC_ID,attr"`
					Source           string `xml:"Source,attr"`
					RelationshipType string `xml:"RelationshipType,attr"`
					Location         struct {
						Text                string `xml:",chardata"`
						CytogeneticLocation string `xml:"CytogeneticLocation"`
						SequenceLocation    []struct {
							Text                     string `xml:",chardata"`
							Assembly                 string `xml:"Assembly,attr"`
							AssemblyAccessionVersion string `xml:"AssemblyAccessionVersion,attr"`
							AssemblyStatus           string `xml:"AssemblyStatus,attr"`
							Chr                      string `xml:"Chr,attr"`
							Accession                string `xml:"Accession,attr"`
							Start                    string `xml:"start,attr"`
							Stop                     string `xml:"stop,attr"`
							DisplayStart             string `xml:"display_start,attr"`
							DisplayStop              string `xml:"display_stop,attr"`
							Strand                   string `xml:"Strand,attr"`
						} `xml:"SequenceLocation"`
					} `xml:"Location"`
					OMIM               string `xml:"OMIM"`
					Haploinsufficiency struct {
						Text          string `xml:",chardata"`
						LastEvaluated string `xml:"last_evaluated,attr"`
						ClinGen       string `xml:"ClinGen,attr"`
					} `xml:"Haploinsufficiency"`
					Triplosensitivity struct {
						Text          string `xml:",chardata"`
						LastEvaluated string `xml:"last_evaluated,attr"`
						ClinGen       string `xml:"ClinGen,attr"`
					} `xml:"Triplosensitivity"`
					Property []string `xml:"Property"`
				} `xml:"Gene"`
			} `xml:"GeneList"`
			Name     string `xml:"Name"`
			Type     string `xml:"Type"`
			Location struct {
				Text                string `xml:",chardata"`
				CytogeneticLocation string `xml:"CytogeneticLocation"`
				SequenceLocation    []struct {
					Text                     string `xml:",chardata"`
					Assembly                 string `xml:"Assembly,attr"`
					AssemblyAccessionVersion string `xml:"AssemblyAccessionVersion,attr"`
					ForDisplay               string `xml:"forDisplay,attr"`
					AssemblyStatus           string `xml:"AssemblyStatus,attr"`
					Chr                      string `xml:"Chr,attr"`
					Accession                string `xml:"Accession,attr"`
					Start                    string `xml:"start,attr"`
					Stop                     string `xml:"stop,attr"`
					DisplayStart             string `xml:"display_start,attr"`
					DisplayStop              string `xml:"display_stop,attr"`
					Length                   string `xml:"Length,attr"`
					PositionVCF              string `xml:"positionVCF,attr"`
					ReferenceAlleleVCF       string `xml:"referenceAlleleVCF,attr"`
					AlternateAlleleVCF       string `xml:"alternateAlleleVCF,attr"`
				} `xml:"SequenceLocation"`
			} `xml:"Location"`
			OtherNameList struct {
				Text string   `xml:",chardata"`
				Name []string `xml:"Name"`
			} `xml:"OtherNameList"`
			XRefList struct {
				Text string `xml:",chardata"`
				XRef []struct {
					Text string `xml:",chardata"`
					Type string `xml:"Type,attr"`
					ID   string `xml:"ID,attr"`
					DB   string `xml:"DB,attr"`
				} `xml:"XRef"`
			} `xml:"XRefList"`
			CanonicalSPDI string `xml:"CanonicalSPDI"`
			HGVSlist      struct {
				Text string `xml:",chardata"`
				HGVS []struct {
					Text                 string `xml:",chardata"`
					Type                 string `xml:"Type,attr"`
					Assembly             string `xml:"Assembly,attr"`
					NucleotideExpression struct {
						Text                     string `xml:",chardata"`
						SequenceAccessionVersion string `xml:"sequenceAccessionVersion,attr"`
						SequenceAccession        string `xml:"sequenceAccession,attr"`
						SequenceVersion          string `xml:"sequenceVersion,attr"`
						Change                   string `xml:"change,attr"`
						Assembly                 string `xml:"Assembly,attr"`
						MANESelect               string `xml:"MANESelect,attr"`
						Expression               string `xml:"Expression"`
					} `xml:"NucleotideExpression"`
					MolecularConsequence struct {
						Text string `xml:",chardata"`
						ID   string `xml:"ID,attr"`
						Type string `xml:"Type,attr"`
						DB   string `xml:"DB,attr"`
					} `xml:"MolecularConsequence"`
					ProteinExpression struct {
						Text                     string `xml:",chardata"`
						SequenceAccessionVersion string `xml:"sequenceAccessionVersion,attr"`
						SequenceAccession        string `xml:"sequenceAccession,attr"`
						SequenceVersion          string `xml:"sequenceVersion,attr"`
						Change                   string `xml:"change,attr"`
						Expression               string `xml:"Expression"`
					} `xml:"ProteinExpression"`
				} `xml:"HGVS"`
			} `xml:"HGVSlist"`
			AlleleFrequencyList struct {
				Text            string `xml:",chardata"`
				AlleleFrequency []struct {
					Text   string `xml:",chardata"`
					Value  string `xml:"Value,attr"`
					Source string `xml:"Source,attr"`
				} `xml:"AlleleFrequency"`
			} `xml:"AlleleFrequencyList"`
			GlobalMinorAlleleFrequency struct {
				Text        string `xml:",chardata"`
				Value       string `xml:"Value,attr"`
				Source      string `xml:"Source,attr"`
				MinorAllele string `xml:"MinorAllele,attr"`
			} `xml:"GlobalMinorAlleleFrequency"`
			ProteinChange []string `xml:"ProteinChange"`
		} `xml:"SimpleAllele"`
		ReviewStatus string `xml:"ReviewStatus"`
		RCVList      struct {
			Text         string `xml:",chardata"`
			RCVAccession []struct {
				Text                     string `xml:",chardata"`
				Title                    string `xml:"Title,attr"`
				DateLastEvaluated        string `xml:"DateLastEvaluated,attr"`
				ReviewStatus             string `xml:"ReviewStatus,attr"`
				Interpretation           string `xml:"Interpretation,attr"`
				SubmissionCount          string `xml:"SubmissionCount,attr"`
				Accession                string `xml:"Accession,attr"`
				Version                  string `xml:"Version,attr"`
				InterpretedConditionList struct {
					Text                 string `xml:",chardata"`
					TraitSetID           string `xml:"TraitSetID,attr"`
					InterpretedCondition struct {
						Text string `xml:",chardata"`
						DB   string `xml:"DB,attr"`
						ID   string `xml:"ID,attr"`
					} `xml:"InterpretedCondition"`
				} `xml:"InterpretedConditionList"`
			} `xml:"RCVAccession"`
		} `xml:"RCVList"`
		Interpretations struct {
			Text           string `xml:",chardata"`
			Interpretation struct {
				Text                string `xml:",chardata"`
				DateLastEvaluated   string `xml:"DateLastEvaluated,attr"`
				NumberOfSubmissions string `xml:"NumberOfSubmissions,attr"`
				NumberOfSubmitters  string `xml:"NumberOfSubmitters,attr"`
				Type                string `xml:"Type,attr"`
				Description         string `xml:"Description"`
				Citation            []struct {
					Text string `xml:",chardata"`
					Type string `xml:"Type,attr"`
					ID   struct {
						Text   string `xml:",chardata"`
						Source string `xml:"Source,attr"`
					} `xml:"ID"`
					URL string `xml:"URL"`
				} `xml:"Citation"`
				ConditionList struct {
					Text     string `xml:",chardata"`
					TraitSet []struct {
						Text  string `xml:",chardata"`
						ID    string `xml:"ID,attr"`
						Type  string `xml:"Type,attr"`
						Trait struct {
							Text string `xml:",chardata"`
							ID   string `xml:"ID,attr"`
							Type string `xml:"Type,attr"`
							Name []struct {
								Text         string `xml:",chardata"`
								ElementValue struct {
									Text string `xml:",chardata"`
									Type string `xml:"Type,attr"`
								} `xml:"ElementValue"`
								XRef []struct {
									Text string `xml:",chardata"`
									Type string `xml:"Type,attr"`
									ID   string `xml:"ID,attr"`
									DB   string `xml:"DB,attr"`
								} `xml:"XRef"`
							} `xml:"Name"`
							XRef []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"ID,attr"`
								DB   string `xml:"DB,attr"`
								Type string `xml:"Type,attr"`
							} `xml:"XRef"`
							Symbol []struct {
								Text         string `xml:",chardata"`
								ElementValue struct {
									Text string `xml:",chardata"`
									Type string `xml:"Type,attr"`
								} `xml:"ElementValue"`
								XRef struct {
									Text string `xml:",chardata"`
									ID   string `xml:"ID,attr"`
									DB   string `xml:"DB,attr"`
									Type string `xml:"Type,attr"`
								} `xml:"XRef"`
							} `xml:"Symbol"`
							AttributeSet []struct {
								Text      string `xml:",chardata"`
								Attribute struct {
									Text         string `xml:",chardata"`
									Type         string `xml:"Type,attr"`
									IntegerValue string `xml:"integerValue,attr"`
								} `xml:"Attribute"`
								XRef struct {
									Text string `xml:",chardata"`
									ID   string `xml:"ID,attr"`
									DB   string `xml:"DB,attr"`
								} `xml:"XRef"`
							} `xml:"AttributeSet"`
							Citation []struct {
								Text   string `xml:",chardata"`
								Type   string `xml:"Type,attr"`
								Abbrev string `xml:"Abbrev,attr"`
								ID     []struct {
									Text   string `xml:",chardata"`
									Source string `xml:"Source,attr"`
								} `xml:"ID"`
							} `xml:"Citation"`
						} `xml:"Trait"`
					} `xml:"TraitSet"`
				} `xml:"ConditionList"`
				DescriptionHistory []struct {
					Text        string `xml:",chardata"`
					Dated       string `xml:"Dated,attr"`
					Description string `xml:"Description"`
				} `xml:"DescriptionHistory"`
			} `xml:"Interpretation"`
		} `xml:"Interpretations"`
		ClinicalAssertionList struct {
			Text              string `xml:",chardata"`
			ClinicalAssertion []struct {
				Text                  string `xml:",chardata"`
				ID                    string `xml:"ID,attr"`
				DateCreated           string `xml:"DateCreated,attr"`
				DateLastUpdated       string `xml:"DateLastUpdated,attr"`
				SubmissionDate        string `xml:"SubmissionDate,attr"`
				FDARecognizedDatabase string `xml:"FDARecognizedDatabase,attr"`
				ClinVarSubmissionID   struct {
					Text                string `xml:",chardata"`
					LocalKey            string `xml:"localKey,attr"`
					Title               string `xml:"title,attr"`
					LocalKeyIsSubmitted string `xml:"localKeyIsSubmitted,attr"`
					SubmittedAssembly   string `xml:"submittedAssembly,attr"`
				} `xml:"ClinVarSubmissionID"`
				ClinVarAccession struct {
					Text                 string `xml:",chardata"`
					Accession            string `xml:"Accession,attr"`
					Type                 string `xml:"Type,attr"`
					Version              string `xml:"Version,attr"`
					SubmitterName        string `xml:"SubmitterName,attr"`
					OrgID                string `xml:"OrgID,attr"`
					OrganizationCategory string `xml:"OrganizationCategory,attr"`
					OrgAbbreviation      string `xml:"OrgAbbreviation,attr"`
				} `xml:"ClinVarAccession"`
				RecordStatus   string `xml:"RecordStatus"`
				ReviewStatus   string `xml:"ReviewStatus"`
				Interpretation struct {
					Text              string `xml:",chardata"`
					DateLastEvaluated string `xml:"DateLastEvaluated,attr"`
					Description       string `xml:"Description"`
					Citation          []struct {
						Text string `xml:",chardata"`
						ID   struct {
							Text   string `xml:",chardata"`
							Source string `xml:"Source,attr"`
						} `xml:"ID"`
						URL string `xml:"URL"`
					} `xml:"Citation"`
					Comment struct {
						Text string `xml:",chardata"`
						Type string `xml:"Type,attr"`
					} `xml:"Comment"`
				} `xml:"Interpretation"`
				Assertion      string `xml:"Assertion"`
				ObservedInList struct {
					Text       string `xml:",chardata"`
					ObservedIn struct {
						Text   string `xml:",chardata"`
						Sample struct {
							Text    string `xml:",chardata"`
							Origin  string `xml:"Origin"`
							Species struct {
								Text       string `xml:",chardata"`
								TaxonomyId string `xml:"TaxonomyId,attr"`
							} `xml:"Species"`
							AffectedStatus string `xml:"AffectedStatus"`
							NumberTested   string `xml:"NumberTested"`
						} `xml:"Sample"`
						Method struct {
							Text         string `xml:",chardata"`
							MethodType   string `xml:"MethodType"`
							TypePlatform string `xml:"TypePlatform"`
						} `xml:"Method"`
						ObservedData struct {
							Text      string `xml:",chardata"`
							Attribute struct {
								Text         string `xml:",chardata"`
								Type         string `xml:"Type,attr"`
								IntegerValue string `xml:"integerValue,attr"`
							} `xml:"Attribute"`
							Citation []struct {
								Text string `xml:",chardata"`
								ID   struct {
									Text   string `xml:",chardata"`
									Source string `xml:"Source,attr"`
								} `xml:"ID"`
							} `xml:"Citation"`
							XRef struct {
								Text string `xml:",chardata"`
								DB   string `xml:"DB,attr"`
								ID   string `xml:"ID,attr"`
								Type string `xml:"Type,attr"`
							} `xml:"XRef"`
						} `xml:"ObservedData"`
					} `xml:"ObservedIn"`
				} `xml:"ObservedInList"`
				SimpleAllele struct {
					Text     string `xml:",chardata"`
					GeneList struct {
						Text string `xml:",chardata"`
						Gene struct {
							Text   string `xml:",chardata"`
							Symbol string `xml:"Symbol,attr"`
						} `xml:"Gene"`
					} `xml:"GeneList"`
					Name          string `xml:"Name"`
					Type          string `xml:"Type"`
					OtherNameList struct {
						Text string `xml:",chardata"`
						Name struct {
							Text string `xml:",chardata"`
							Type string `xml:"Type,attr"`
						} `xml:"Name"`
					} `xml:"OtherNameList"`
					XRefList struct {
						Text string `xml:",chardata"`
						XRef struct {
							Text string `xml:",chardata"`
							DB   string `xml:"DB,attr"`
							ID   string `xml:"ID,attr"`
							Type string `xml:"Type,attr"`
						} `xml:"XRef"`
					} `xml:"XRefList"`
					AttributeSet struct {
						Text      string `xml:",chardata"`
						Attribute struct {
							Text string `xml:",chardata"`
							Type string `xml:"Type,attr"`
						} `xml:"Attribute"`
					} `xml:"AttributeSet"`
					Location struct {
						Text             string `xml:",chardata"`
						SequenceLocation struct {
							Text            string `xml:",chardata"`
							Assembly        string `xml:"Assembly,attr"`
							Chr             string `xml:"Chr,attr"`
							AlternateAllele string `xml:"alternateAllele,attr"`
							ReferenceAllele string `xml:"referenceAllele,attr"`
							Start           string `xml:"start,attr"`
							Stop            string `xml:"stop,attr"`
							Length          string `xml:"Length,attr"`
						} `xml:"SequenceLocation"`
					} `xml:"Location"`
				} `xml:"SimpleAllele"`
				TraitSet struct {
					Text  string `xml:",chardata"`
					Type  string `xml:"Type,attr"`
					Trait struct {
						Text string `xml:",chardata"`
						Type string `xml:"Type,attr"`
						Name struct {
							Text         string `xml:",chardata"`
							ElementValue struct {
								Text string `xml:",chardata"`
								Type string `xml:"Type,attr"`
							} `xml:"ElementValue"`
						} `xml:"Name"`
						XRef struct {
							Text string `xml:",chardata"`
							DB   string `xml:"DB,attr"`
							ID   string `xml:"ID,attr"`
							Type string `xml:"Type,attr"`
						} `xml:"XRef"`
					} `xml:"Trait"`
				} `xml:"TraitSet"`
				AttributeSet []struct {
					Text      string `xml:",chardata"`
					Attribute struct {
						Text string `xml:",chardata"`
						Type string `xml:"Type,attr"`
					} `xml:"Attribute"`
					Citation struct {
						Text string `xml:",chardata"`
						URL  string `xml:"URL"`
						ID   struct {
							Text   string `xml:",chardata"`
							Source string `xml:"Source,attr"`
						} `xml:"ID"`
					} `xml:"Citation"`
				} `xml:"AttributeSet"`
				SubmissionNameList struct {
					Text           string `xml:",chardata"`
					SubmissionName string `xml:"SubmissionName"`
				} `xml:"SubmissionNameList"`
				Comment string `xml:"Comment"`
			} `xml:"ClinicalAssertion"`
		} `xml:"ClinicalAssertionList"`
		TraitMappingList struct {
			Text         string `xml:",chardata"`
			TraitMapping []struct {
				Text                string `xml:",chardata"`
				ClinicalAssertionID string `xml:"ClinicalAssertionID,attr"`
				TraitType           string `xml:"TraitType,attr"`
				MappingType         string `xml:"MappingType,attr"`
				MappingValue        string `xml:"MappingValue,attr"`
				MappingRef          string `xml:"MappingRef,attr"`
				MedGen              struct {
					Text string `xml:",chardata"`
					CUI  string `xml:"CUI,attr"`
					Name string `xml:"Name,attr"`
				} `xml:"MedGen"`
			} `xml:"TraitMapping"`
		} `xml:"TraitMappingList"`
	} `xml:"InterpretedRecord"`
}

type ClinVarDataReleaseInfo struct {
	W3SchemaInfo         string
	ClinVarSchemaVersion string
	ClinVarReleaseDate   string
}

type ClinVarVariationData struct {
	Accesssion              string
	Version                 string
	Type                    string
	GeneAffected            string
	GeneEntrezID            string
	GeneOmimID              string
	NcbiRefSeq              string
	LocationType            string
	DbSNPID                 string
	GenomeVersion           string
	ChromLocation           string
	ChromStart              string
	ChromStop               string
	Length                  string
	OmimID                  string
	ReviewStatus            string
	HGVData                 []HGVData
	RCVData                 []RCVData
	ClinicalInterpretations ClinicalInterpretations
	// ClinicalAssertions  []ClinicalAssertions
}

type HGVData struct {
	Consequence string
}

type RCVData struct {
	AccessionID     string
	Version         string
	Interpretation  string
	Condition       string
	SubmissionCount string
	ReviewStatus    string
	MedGenID        string
	TraitSetID      string
}

type ClinicalInterpretations struct {
	Citations []Citations
	Trait     []Traits
}

type Citations struct {
	CitationSource string
	CitationID     string
}

type Traits struct {
	ID               string
	Name             string
	Citations        []Citations
	PhenotypicSeries string
	MIM              string
	MedGen           string
	Orph             string
}

// type ClinicalAssertions struct {
// 	ClinicalAssertionAccession         string
// 	ClinicalAssertionVersion           string
// 	ClinicalAssertionID                string
// 	ClinicalAssertionSubmitterName     string
// 	ClinicalAssertionInterpretation    string
// 	ClinicalAssertionPhenoSeries       string
// 	ClinicalAssertionMESH              string
// 	ClinicalAssertionOrphanet          string
// 	ClinicalAssertionMedgen            string
// 	ClinicalAssertionMethodType        string
// 	ClinicalAssertionSCVID             string
// 	ClinicalAssertionSCVIDReviewStatus string
// 	TraitMappingDisease                string
// }

func main() {
	//Define default flag values and enable input from command line
	inputXML := flag.String("i", "", "Path of XML file to open")
	releaseData := flag.String("r", "", "Indication if only the ClinVar release schemas and data are desired")
	outputFile := flag.String("o", "", "Path of file to write")
	flag.Parse()

	data, err := parseXMLFileToData(*inputXML)
	if err != nil {
		log.Fatal("Could not parse XML file", err)
	}

	//Obtain top-level info for ClinVar file being used
	if *releaseData == "yes" {
		var releaseInfo ClinVarDataReleaseInfo
		releaseInfo = data.extractReleaseInfo()
		spew.Dump(releaseInfo)
	}

	//Obtain top-level information for variants
	var allVariantsData []ClinVarVariationData
	allVariantsData = data.extractAllVariants()

	jsonMarshal, err := json.Marshal(allVariantsData)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else if *outputFile == "" {
		fmt.Println(string(jsonMarshal))
		return
	} else {
		writeClinVarVariationDataFile(*outputFile, jsonMarshal)
	}

}

func parseXMLFileToData(file string) (*ClinVarDataRelease, error) {
	variantFile := os.Stdin
	var err error
	// Open XML file for reading
	if len(file) > 0 {
		variantFile, err = os.Open(file)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Fprintln(os.Stderr, "Variant info file has opened successfully!")
		}

	}
	defer variantFile.Close()

	// Create XML decoder
	xmlDecoder := xml.NewDecoder(variantFile)

	//Parse XML file into the variant data struct variable (top level ClinVarDataRelease wrapper)
	var data ClinVarDataRelease
	err = xmlDecoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (data ClinVarDataRelease) extractReleaseInfo() ClinVarDataReleaseInfo {
	clinRelease := ClinVarDataReleaseInfo{}
	clinRelease.W3SchemaInfo = data.Xsi
	clinRelease.ClinVarSchemaVersion = data.NoNamespaceSchemaLocation
	clinRelease.ClinVarReleaseDate = data.ReleaseDate
	return clinRelease
}

func (data *ClinVarDataRelease) extractAllVariants() []ClinVarVariationData {
	allVariantsInfo := []ClinVarVariationData{}
	for _, variant := range data.Variants {
		singleVariantInfo := variant.extractClinVarVariantData()
		allVariantsInfo = append(allVariantsInfo, singleVariantInfo)
	}
	return allVariantsInfo
}

func (variant *VariationArchive) extractClinVarVariantData() ClinVarVariationData {

	singleVariantInfo := ClinVarVariationData{}

	singleVariantInfo.Accesssion = variant.Accession
	singleVariantInfo.Version = variant.Version
	singleVariantInfo.Type = variant.VariationType

	if len(variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation) > 0 && len(variant.InterpretedRecord.SimpleAllele.GeneList.Gene) > 0 && len(variant.InterpretedRecord.SimpleAllele.GeneList.Gene[0].Location.SequenceLocation) > 0 {
		if variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation[0].Start >= variant.InterpretedRecord.SimpleAllele.GeneList.Gene[0].Location.SequenceLocation[0].Start && variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation[0].Stop <= variant.InterpretedRecord.SimpleAllele.GeneList.Gene[0].Location.SequenceLocation[0].Stop {
			singleVariantInfo.LocationType = "withinGene"
		} else {
			singleVariantInfo.LocationType = "outsideGene"
		}
	} else {
		singleVariantInfo.LocationType = "notProvided"
	}

	if len(variant.InterpretedRecord.SimpleAllele.GeneList.Gene) > 0 {
		singleVariantInfo.GeneAffected = variant.InterpretedRecord.SimpleAllele.GeneList.Gene[0].Symbol
		singleVariantInfo.GeneEntrezID = variant.InterpretedRecord.SimpleAllele.GeneList.Gene[0].GeneID
		singleVariantInfo.GeneOmimID = variant.InterpretedRecord.SimpleAllele.GeneList.Gene[0].OMIM
		//TODO: spew.Dump()
	} else {
		singleVariantInfo.GeneAffected = "notProvided"
		singleVariantInfo.GeneEntrezID = "notProvided"
		singleVariantInfo.GeneOmimID = "notProvided"
	}

	singleVariantInfo.NcbiRefSeq = variant.InterpretedRecord.SimpleAllele.CanonicalSPDI

	if len(variant.InterpretedRecord.SimpleAllele.XRefList.XRef) > 0 {
		for _, snpdbXref := range variant.InterpretedRecord.SimpleAllele.XRefList.XRef {
			if snpdbXref.DB == "dbSNP" {
				singleVariantInfo.DbSNPID = snpdbXref.ID
			} else {
				singleVariantInfo.DbSNPID = "notProvided"
			}
		}
	}

	singleVariantInfo.ChromLocation = variant.InterpretedRecord.SimpleAllele.Location.CytogeneticLocation

	if len(variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation) > 0 {
		singleVariantInfo.GenomeVersion = variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation[0].Assembly
		singleVariantInfo.ChromStart = variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation[0].Start
		singleVariantInfo.ChromStop = variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation[0].Stop
		singleVariantInfo.Length = variant.InterpretedRecord.SimpleAllele.Location.SequenceLocation[0].Length
	} else {
		singleVariantInfo.GenomeVersion = "notProvided"
		singleVariantInfo.ChromStart = "notProvided"
		singleVariantInfo.ChromStop = "notProvided"
		singleVariantInfo.Length = "notProvided"
	}

	if len(variant.InterpretedRecord.SimpleAllele.XRefList.XRef) > 0 {
		for _, omimXref := range variant.InterpretedRecord.SimpleAllele.XRefList.XRef {
			if omimXref.DB == "OMIM" {
				singleVariantInfo.OmimID = omimXref.ID
			} else {
				singleVariantInfo.OmimID = "notProvided"
			}
		}
	}

	singleVariantInfo.ReviewStatus = variant.InterpretedRecord.ReviewStatus

	variantHgvConsequence := []HGVData{}
	for _, hgvs := range variant.InterpretedRecord.SimpleAllele.HGVSlist.HGVS {
		nc := ""
		if hgvs.MolecularConsequence.Type != "" {
			nc = hgvs.MolecularConsequence.Type
			skip := false
			for _, cons := range variantHgvConsequence {
				oc := cons.Consequence
				if nc == oc {
					skip = true
					break
				}
			}
			if !skip {
				variantHgvConsequence = append(variantHgvConsequence, HGVData{
					Consequence: nc})
			}
		}
	}
	singleVariantInfo.HGVData = variantHgvConsequence

	variantAllRcvs := []RCVData{}
	for _, rcvs := range variant.InterpretedRecord.RCVList.RCVAccession {
		if rcvs.InterpretedConditionList.InterpretedCondition.DB == "MedGen" {
			variantAllRcvs = append(variantAllRcvs, RCVData{
				AccessionID:     rcvs.Accession,
				Version:         rcvs.Version,
				Interpretation:  rcvs.Interpretation,
				Condition:       rcvs.InterpretedConditionList.InterpretedCondition.Text,
				SubmissionCount: rcvs.SubmissionCount,
				ReviewStatus:    rcvs.ReviewStatus,
				MedGenID:        rcvs.InterpretedConditionList.InterpretedCondition.ID,
				TraitSetID:      rcvs.InterpretedConditionList.TraitSetID})

		} else {
			variantAllRcvs = append(variantAllRcvs, RCVData{
				AccessionID:     rcvs.Accession,
				Version:         rcvs.Version,
				Interpretation:  rcvs.Interpretation,
				SubmissionCount: rcvs.SubmissionCount,
				ReviewStatus:    rcvs.ReviewStatus,
				MedGenID:        "notProvided",
				Condition:       "notProvided",
				TraitSetID:      rcvs.InterpretedConditionList.TraitSetID})
		}

	}
	singleVariantInfo.RCVData = variantAllRcvs

	variantAllCitations := []Citations{}
	for _, citations := range variant.InterpretedRecord.Interpretations.Interpretation.Citation {
		variantAllCitations = append(variantAllCitations, Citations{
			CitationSource: citations.ID.Source,
			CitationID:     citations.ID.Text})
	}
	singleVariantInfo.ClinicalInterpretations.Citations = variantAllCitations

	variantAllTraits := []Traits{}
	for _, trait := range variant.InterpretedRecord.Interpretations.Interpretation.ConditionList.TraitSet {
		for _, name := range trait.Trait.Name {
			nt := Traits{ID: trait.Trait.ID}

			if name.ElementValue.Type == "Preferred" {
				nt.Name = name.ElementValue.Text
			} else {
				continue
			}
			for _, xRefTrait := range trait.Trait.XRef {
				if xRefTrait.DB == "Orphanet" {
					nt.Orph = xRefTrait.ID
				} else if xRefTrait.DB == "MedGen" {
					nt.MedGen = xRefTrait.ID
				} else if xRefTrait.DB == "OMIM" {
					if xRefTrait.Type == "Phenotypic series" {
						nt.PhenotypicSeries = xRefTrait.ID
					} else if xRefTrait.Type == "MIM" {
						nt.MIM = xRefTrait.ID
					}
				}
			}
			for _, citations := range trait.Trait.Citation {
				for _, citationsInfo := range citations.ID {
					nt.Citations = append(nt.Citations, Citations{
						CitationSource: citationsInfo.Source,
						CitationID:     citationsInfo.Text})
				}
			}
			variantAllTraits = append(variantAllTraits, nt)
		}
	}
	singleVariantInfo.ClinicalInterpretations.Trait = variantAllTraits

	return singleVariantInfo
}

func writeClinVarVariationDataFile(outputFile string, jsonMarshal []byte) {
	out := os.Stdout
	var err error
	if len(outputFile) > 0 {
		out, err = os.Create(outputFile)
		if err != nil {
			log.Fatal("Could not create outputfile: ", outputFile, "\n", err)
		}
	}
	defer out.Close()
	out.WriteString(string(jsonMarshal))
}
