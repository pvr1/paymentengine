// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:reda.071.001.01
package reda_071_001_01

import (
	"encoding/xml"
)

// Element
type Document struct {
	XMLName xml.Name `xml:"Document"`

	ReqToPayDbtrActvtnAmdmntReq RequestToPayDebtorActivationAmendmentRequestV01 `xml:",any"`
}

// XSD ComplexType declarations

type ActivationHeader2 struct {
	XMLName xml.Name

	MsgId Max35Text `xml:"MsgId"`

	CreDtTm ISODateTime `xml:"CreDtTm"`

	MsgOrgtr *RTPPartyIdentification1 `xml:"MsgOrgtr"`

	MsgRcpt *RTPPartyIdentification1 `xml:"MsgRcpt"`

	InitgPty RTPPartyIdentification1 `xml:"InitgPty"`

	InnerXml string `xml:",innerxml"`
}

type AddressType3Choice struct {
	XMLName xml.Name

	Cd *AddressType2Code `xml:"Cd"`

	Prtry *GenericIdentification30 `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type Contact4 struct {
	XMLName xml.Name

	NmPrfx *NamePrefix2Code `xml:"NmPrfx"`

	Nm *Max140Text `xml:"Nm"`

	PhneNb *PhoneNumber `xml:"PhneNb"`

	MobNb *PhoneNumber `xml:"MobNb"`

	FaxNb *PhoneNumber `xml:"FaxNb"`

	EmailAdr *Max2048Text `xml:"EmailAdr"`

	EmailPurp *Max35Text `xml:"EmailPurp"`

	JobTitl *Max35Text `xml:"JobTitl"`

	Rspnsblty *Max35Text `xml:"Rspnsblty"`

	Dept *Max70Text `xml:"Dept"`

	Othr []OtherContact1 `xml:"Othr"`

	PrefrdMtd *PreferredContactMethod1Code `xml:"PrefrdMtd"`

	InnerXml string `xml:",innerxml"`
}

type ContractReference1 struct {
	XMLName xml.Name

	Tp *DocumentType1Choice `xml:"Tp"`

	Ref Max500Text `xml:"Ref"`

	InnerXml string `xml:",innerxml"`
}

type DateAndDateTime2Choice struct {
	XMLName xml.Name

	Dt *ISODate `xml:"Dt"`

	DtTm *ISODateTime `xml:"DtTm"`

	InnerXml string `xml:",innerxml"`
}

type DateAndPlaceOfBirth1 struct {
	XMLName xml.Name

	BirthDt ISODate `xml:"BirthDt"`

	PrvcOfBirth *Max35Text `xml:"PrvcOfBirth"`

	CityOfBirth Max35Text `xml:"CityOfBirth"`

	CtryOfBirth CountryCode `xml:"CtryOfBirth"`

	InnerXml string `xml:",innerxml"`
}

type DebtorActivation3 struct {
	XMLName xml.Name

	DbtrActvtnId *Max35Text `xml:"DbtrActvtnId"`

	DispNm *Max140Text `xml:"DispNm"`

	UltmtDbtr *RTPPartyIdentification1 `xml:"UltmtDbtr"`

	Dbtr RTPPartyIdentification1 `xml:"Dbtr"`

	DbtrSolPrvdr RTPPartyIdentification1 `xml:"DbtrSolPrvdr"`

	CstmrId []Party49Choice `xml:"CstmrId"`

	CtrctFrmtTp []DocumentFormat2Choice `xml:"CtrctFrmtTp"`

	CtrctRef []ContractReference1 `xml:"CtrctRef"`

	Cdtr RTPPartyIdentification1 `xml:"Cdtr"`

	UltmtCdtr *RTPPartyIdentification1 `xml:"UltmtCdtr"`

	ActvtnReqDlvryPty *RTPPartyIdentification1 `xml:"ActvtnReqDlvryPty"`

	StartDt *DateAndDateTime2Choice `xml:"StartDt"`

	EndDt *DateAndDateTime2Choice `xml:"EndDt"`

	DdctdActvtnCd *Max35Text `xml:"DdctdActvtnCd"`

	InnerXml string `xml:",innerxml"`
}

type DebtorActivation4 struct {
	XMLName xml.Name

	DbtrActvtnId *Max35Text `xml:"DbtrActvtnId"`

	DispNm *Max140Text `xml:"DispNm"`

	UltmtDbtr *RTPPartyIdentification1 `xml:"UltmtDbtr"`

	Dbtr *RTPPartyIdentification1 `xml:"Dbtr"`

	DbtrSolPrvdr *RTPPartyIdentification1 `xml:"DbtrSolPrvdr"`

	CstmrId []Party49Choice `xml:"CstmrId"`

	CtrctFrmtTp []DocumentFormat2Choice `xml:"CtrctFrmtTp"`

	CtrctRef []ContractReference1 `xml:"CtrctRef"`

	Cdtr *RTPPartyIdentification1 `xml:"Cdtr"`

	UltmtCdtr *RTPPartyIdentification1 `xml:"UltmtCdtr"`

	ActvtnReqDlvryPty *RTPPartyIdentification1 `xml:"ActvtnReqDlvryPty"`

	StartDt *DateAndDateTime2Choice `xml:"StartDt"`

	EndDt *DateAndDateTime2Choice `xml:"EndDt"`

	DdctdActvtnCd *Max35Text `xml:"DdctdActvtnCd"`

	InnerXml string `xml:",innerxml"`
}

type DebtorActivationAmendment3 struct {
	XMLName xml.Name

	OrgnlBizInstr *OriginalBusinessInstruction1 `xml:"OrgnlBizInstr"`

	AmdmntRsn *DebtorActivationAmendmentReason2 `xml:"AmdmntRsn"`

	Amdmnt DebtorActivationAmendment4 `xml:"Amdmnt"`

	OrgnlActvtn OriginalActivation2Choice `xml:"OrgnlActvtn"`

	SplmtryData []SupplementaryData1 `xml:"SplmtryData"`

	InnerXml string `xml:",innerxml"`
}

type DebtorActivationAmendment4 struct {
	XMLName xml.Name

	DbtrActvtn *DebtorActivation4 `xml:"DbtrActvtn"`

	ElctrncInvcData *ElectronicInvoice1 `xml:"ElctrncInvcData"`

	InnerXml string `xml:",innerxml"`
}

type DebtorActivationAmendmentReason1Choice struct {
	XMLName xml.Name

	Cd *ExternalDebtorActivationAmendmentReason1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type DebtorActivationAmendmentReason2 struct {
	XMLName xml.Name

	Orgtr *RTPPartyIdentification1 `xml:"Orgtr"`

	Rsn DebtorActivationAmendmentReason1Choice `xml:"Rsn"`

	AddtlInf []Max105Text `xml:"AddtlInf"`

	InnerXml string `xml:",innerxml"`
}

type DocumentFormat2Choice struct {
	XMLName xml.Name

	Cd *ExternalDocumentFormat1Code `xml:"Cd"`

	Prtry *GenericIdentification1 `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type DocumentType1Choice struct {
	XMLName xml.Name

	Cd *ExternalDocumentType1Code `xml:"Cd"`

	Prtry *GenericIdentification1 `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ElectronicInvoice1 struct {
	XMLName xml.Name

	PresntmntTp PresentmentType1Code `xml:",any"`

	InnerXml string `xml:",innerxml"`
}

type GenericIdentification1 struct {
	XMLName xml.Name

	Id Max35Text `xml:"Id"`

	SchmeNm *Max35Text `xml:"SchmeNm"`

	Issr *Max35Text `xml:"Issr"`

	InnerXml string `xml:",innerxml"`
}

type GenericIdentification30 struct {
	XMLName xml.Name

	Id Exact4AlphaNumericText `xml:"Id"`

	Issr Max35Text `xml:"Issr"`

	SchmeNm *Max35Text `xml:"SchmeNm"`

	InnerXml string `xml:",innerxml"`
}

type GenericOrganisationIdentification1 struct {
	XMLName xml.Name

	Id Max35Text `xml:"Id"`

	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm"`

	Issr *Max35Text `xml:"Issr"`

	InnerXml string `xml:",innerxml"`
}

type GenericPersonIdentification1 struct {
	XMLName xml.Name

	Id Max35Text `xml:"Id"`

	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm"`

	Issr *Max35Text `xml:"Issr"`

	InnerXml string `xml:",innerxml"`
}

type OrganisationIdentification37 struct {
	XMLName xml.Name

	AnyBIC *AnyBICDec2014Identifier `xml:"AnyBIC"`

	LEI *LEIIdentifier `xml:"LEI"`

	EmailAdr *Max256Text `xml:"EmailAdr"`

	Othr []GenericOrganisationIdentification1 `xml:"Othr"`

	InnerXml string `xml:",innerxml"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	XMLName xml.Name

	Cd *ExternalOrganisationIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type OriginalActivation2Choice struct {
	XMLName xml.Name

	OrgnlDbtrId *Party49Choice `xml:"OrgnlDbtrId"`

	OrgnlActvtnData *DebtorActivation3 `xml:"OrgnlActvtnData"`

	InnerXml string `xml:",innerxml"`
}

type OriginalBusinessInstruction1 struct {
	XMLName xml.Name

	MsgId Max35Text `xml:"MsgId"`

	MsgNmId *Max35Text `xml:"MsgNmId"`

	CreDtTm *ISODateTime `xml:"CreDtTm"`

	InnerXml string `xml:",innerxml"`
}

type OtherContact1 struct {
	XMLName xml.Name

	ChanlTp Max4Text `xml:"ChanlTp"`

	Id *Max128Text `xml:"Id"`

	InnerXml string `xml:",innerxml"`
}

type Party49Choice struct {
	XMLName xml.Name

	OrgId *OrganisationIdentification37 `xml:"OrgId"`

	PrvtId *PersonIdentification17 `xml:"PrvtId"`

	InnerXml string `xml:",innerxml"`
}

type PersonIdentification17 struct {
	XMLName xml.Name

	DtAndPlcOfBirth *DateAndPlaceOfBirth1 `xml:"DtAndPlcOfBirth"`

	EmailAdr *Max256Text `xml:"EmailAdr"`

	Othr []GenericPersonIdentification1 `xml:"Othr"`

	InnerXml string `xml:",innerxml"`
}

type PersonIdentificationSchemeName1Choice struct {
	XMLName xml.Name

	Cd *ExternalPersonIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type PostalAddress24 struct {
	XMLName xml.Name

	AdrTp *AddressType3Choice `xml:"AdrTp"`

	Dept *Max70Text `xml:"Dept"`

	SubDept *Max70Text `xml:"SubDept"`

	StrtNm *Max70Text `xml:"StrtNm"`

	BldgNb *Max16Text `xml:"BldgNb"`

	BldgNm *Max35Text `xml:"BldgNm"`

	Flr *Max70Text `xml:"Flr"`

	PstBx *Max16Text `xml:"PstBx"`

	Room *Max70Text `xml:"Room"`

	PstCd *Max16Text `xml:"PstCd"`

	TwnNm *Max35Text `xml:"TwnNm"`

	TwnLctnNm *Max35Text `xml:"TwnLctnNm"`

	DstrctNm *Max35Text `xml:"DstrctNm"`

	CtrySubDvsn *Max35Text `xml:"CtrySubDvsn"`

	Ctry *CountryCode `xml:"Ctry"`

	AdrLine []Max70Text `xml:"AdrLine"`

	InnerXml string `xml:",innerxml"`
}

type RTPPartyIdentification1 struct {
	XMLName xml.Name

	Nm *Max140Text `xml:"Nm"`

	PstlAdr *PostalAddress24 `xml:"PstlAdr"`

	Id *Party49Choice `xml:"Id"`

	CtryOfRes *CountryCode `xml:"CtryOfRes"`

	CtctDtls *Contact4 `xml:"CtctDtls"`

	InnerXml string `xml:",innerxml"`
}

type RequestToPayDebtorActivationAmendmentRequestV01 struct {
	XMLName xml.Name

	Hdr ActivationHeader2 `xml:"Hdr"`

	AmdmntData []DebtorActivationAmendment3 `xml:"AmdmntData"`

	SplmtryData []SupplementaryData1 `xml:"SplmtryData"`

	InnerXml string `xml:",innerxml"`
}

type SupplementaryData1 struct {
	XMLName xml.Name

	PlcAndNm *Max350Text `xml:"PlcAndNm"`

	Envlp SupplementaryDataEnvelope1 `xml:"Envlp"`

	InnerXml string `xml:",innerxml"`
}

type SupplementaryDataEnvelope1 struct {
	XMLName xml.Name

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type AddressType2Code string

const AddressType2CodeAddr AddressType2Code = "ADDR"

const AddressType2CodePbox AddressType2Code = "PBOX"

const AddressType2CodeHome AddressType2Code = "HOME"

const AddressType2CodeBizz AddressType2Code = "BIZZ"

const AddressType2CodeMlto AddressType2Code = "MLTO"

const AddressType2CodeDlvy AddressType2Code = "DLVY"

type AnyBICDec2014Identifier string

type CountryCode string

type Exact4AlphaNumericText string

type ExternalDebtorActivationAmendmentReason1Code string

type ExternalDocumentFormat1Code string

type ExternalDocumentType1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type ISODate string

type ISODateTime string

type LEIIdentifier string

type Max105Text string

type Max128Text string

type Max140Text string

type Max16Text string

type Max2048Text string

type Max256Text string

type Max350Text string

type Max35Text string

type Max4Text string

type Max500Text string

type Max70Text string

type NamePrefix2Code string

const NamePrefix2CodeDoct NamePrefix2Code = "DOCT"

const NamePrefix2CodeMadm NamePrefix2Code = "MADM"

const NamePrefix2CodeMiss NamePrefix2Code = "MISS"

const NamePrefix2CodeMist NamePrefix2Code = "MIST"

const NamePrefix2CodeMiks NamePrefix2Code = "MIKS"

type PhoneNumber string

type PreferredContactMethod1Code string

const PreferredContactMethod1CodeLett PreferredContactMethod1Code = "LETT"

const PreferredContactMethod1CodeMail PreferredContactMethod1Code = "MAIL"

const PreferredContactMethod1CodePhon PreferredContactMethod1Code = "PHON"

const PreferredContactMethod1CodeFaxx PreferredContactMethod1Code = "FAXX"

const PreferredContactMethod1CodeCell PreferredContactMethod1Code = "CELL"

type PresentmentType1Code string

const PresentmentType1CodeFull PresentmentType1Code = "FULL"

const PresentmentType1CodePayd PresentmentType1Code = "PAYD"
