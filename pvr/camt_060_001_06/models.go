// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.060.001.06
package camt_060_001_06

import (
	"encoding/xml"
)

// Element
type Document struct {
	XMLName xml.Name `xml:"Document"`

	AcctRptgReq AccountReportingRequestV06 `xml:",any"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	XMLName xml.Name

	IBAN *IBAN2007Identifier `xml:"IBAN"`

	Othr *GenericAccountIdentification1 `xml:"Othr"`

	InnerXml string `xml:",innerxml"`
}

type AccountReportingRequestV06 struct {
	XMLName xml.Name

	GrpHdr GroupHeader77 `xml:"GrpHdr"`

	RptgReq []ReportingRequest6 `xml:"RptgReq"`

	SplmtryData []SupplementaryData1 `xml:"SplmtryData"`

	InnerXml string `xml:",innerxml"`
}

type AccountSchemeName1Choice struct {
	XMLName xml.Name

	Cd *ExternalAccountIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	XMLName xml.Name

	Ccy ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`

	Text     string `xml:",chardata"`
	InnerXml string `xml:",innerxml"`
}

type AddressType3Choice struct {
	XMLName xml.Name

	Cd *AddressType2Code `xml:"Cd"`

	Prtry *GenericIdentification30 `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type BalanceSubType1Choice struct {
	XMLName xml.Name

	Cd *ExternalBalanceSubType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type BalanceType10Choice struct {
	XMLName xml.Name

	Cd *ExternalBalanceType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type BalanceType13 struct {
	XMLName xml.Name

	CdOrPrtry BalanceType10Choice `xml:"CdOrPrtry"`

	SubTp *BalanceSubType1Choice `xml:"SubTp"`

	InnerXml string `xml:",innerxml"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	XMLName xml.Name

	FinInstnId FinancialInstitutionIdentification18 `xml:"FinInstnId"`

	BrnchId *BranchData3 `xml:"BrnchId"`

	InnerXml string `xml:",innerxml"`
}

type BranchData3 struct {
	XMLName xml.Name

	Id *Max35Text `xml:"Id"`

	LEI *LEIIdentifier `xml:"LEI"`

	Nm *Max140Text `xml:"Nm"`

	PstlAdr *PostalAddress24 `xml:"PstlAdr"`

	InnerXml string `xml:",innerxml"`
}

type CashAccount40 struct {
	XMLName xml.Name

	Id *AccountIdentification4Choice `xml:"Id"`

	Tp *CashAccountType2Choice `xml:"Tp"`

	Ccy *ActiveOrHistoricCurrencyCode `xml:"Ccy"`

	Nm *Max70Text `xml:"Nm"`

	Prxy *ProxyAccountIdentification1 `xml:"Prxy"`

	InnerXml string `xml:",innerxml"`
}

type CashAccountType2Choice struct {
	XMLName xml.Name

	Cd *ExternalCashAccountType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ClearingSystemIdentification2Choice struct {
	XMLName xml.Name

	Cd *ExternalClearingSystemIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ClearingSystemMemberIdentification2 struct {
	XMLName xml.Name

	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId"`

	MmbId Max35Text `xml:"MmbId"`

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

type DateAndPlaceOfBirth1 struct {
	XMLName xml.Name

	BirthDt ISODate `xml:"BirthDt"`

	PrvcOfBirth *Max35Text `xml:"PrvcOfBirth"`

	CityOfBirth Max35Text `xml:"CityOfBirth"`

	CtryOfBirth CountryCode `xml:"CtryOfBirth"`

	InnerXml string `xml:",innerxml"`
}

type DatePeriod3 struct {
	XMLName xml.Name

	FrDt ISODate `xml:"FrDt"`

	ToDt *ISODate `xml:"ToDt"`

	InnerXml string `xml:",innerxml"`
}

type EntryStatus1Choice struct {
	XMLName xml.Name

	Cd *ExternalEntryStatus1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type FinancialIdentificationSchemeName1Choice struct {
	XMLName xml.Name

	Cd *ExternalFinancialInstitutionIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type FinancialInstitutionIdentification18 struct {
	XMLName xml.Name

	BICFI *BICFIDec2014Identifier `xml:"BICFI"`

	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`

	LEI *LEIIdentifier `xml:"LEI"`

	Nm *Max140Text `xml:"Nm"`

	PstlAdr *PostalAddress24 `xml:"PstlAdr"`

	Othr *GenericFinancialIdentification1 `xml:"Othr"`

	InnerXml string `xml:",innerxml"`
}

type GenericAccountIdentification1 struct {
	XMLName xml.Name

	Id Max34Text `xml:"Id"`

	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm"`

	Issr *Max35Text `xml:"Issr"`

	InnerXml string `xml:",innerxml"`
}

type GenericFinancialIdentification1 struct {
	XMLName xml.Name

	Id Max35Text `xml:"Id"`

	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm"`

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

type GroupHeader77 struct {
	XMLName xml.Name

	MsgId Max35Text `xml:"MsgId"`

	CreDtTm ISODateTime `xml:"CreDtTm"`

	MsgSndr *Party40Choice `xml:"MsgSndr"`

	InnerXml string `xml:",innerxml"`
}

type Limit2 struct {
	XMLName xml.Name

	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	CdtDbtInd FloorLimitType1Code `xml:"CdtDbtInd"`

	InnerXml string `xml:",innerxml"`
}

type OrganisationIdentification29 struct {
	XMLName xml.Name

	AnyBIC *AnyBICDec2014Identifier `xml:"AnyBIC"`

	LEI *LEIIdentifier `xml:"LEI"`

	Othr []GenericOrganisationIdentification1 `xml:"Othr"`

	InnerXml string `xml:",innerxml"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	XMLName xml.Name

	Cd *ExternalOrganisationIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type OtherContact1 struct {
	XMLName xml.Name

	ChanlTp Max4Text `xml:"ChanlTp"`

	Id *Max128Text `xml:"Id"`

	InnerXml string `xml:",innerxml"`
}

type Party38Choice struct {
	XMLName xml.Name

	OrgId *OrganisationIdentification29 `xml:"OrgId"`

	PrvtId *PersonIdentification13 `xml:"PrvtId"`

	InnerXml string `xml:",innerxml"`
}

type Party40Choice struct {
	XMLName xml.Name

	Pty *PartyIdentification135 `xml:"Pty"`

	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`

	InnerXml string `xml:",innerxml"`
}

type PartyIdentification135 struct {
	XMLName xml.Name

	Nm *Max140Text `xml:"Nm"`

	PstlAdr *PostalAddress24 `xml:"PstlAdr"`

	Id *Party38Choice `xml:"Id"`

	CtryOfRes *CountryCode `xml:"CtryOfRes"`

	CtctDtls *Contact4 `xml:"CtctDtls"`

	InnerXml string `xml:",innerxml"`
}

type PersonIdentification13 struct {
	XMLName xml.Name

	DtAndPlcOfBirth *DateAndPlaceOfBirth1 `xml:"DtAndPlcOfBirth"`

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

type ProxyAccountIdentification1 struct {
	XMLName xml.Name

	Tp *ProxyAccountType1Choice `xml:"Tp"`

	Id Max2048Text `xml:"Id"`

	InnerXml string `xml:",innerxml"`
}

type ProxyAccountType1Choice struct {
	XMLName xml.Name

	Cd *ExternalProxyAccountType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ReportingPeriod5 struct {
	XMLName xml.Name

	FrToDt DatePeriod3 `xml:"FrToDt"`

	FrToTm *TimePeriodDetails1 `xml:"FrToTm"`

	Tp QueryType3Code `xml:"Tp"`

	InnerXml string `xml:",innerxml"`
}

type ReportingRequest6 struct {
	XMLName xml.Name

	Id *Max35Text `xml:"Id"`

	ReqdMsgNmId Max35Text `xml:"ReqdMsgNmId"`

	Acct *CashAccount40 `xml:"Acct"`

	AcctOwnr Party40Choice `xml:"AcctOwnr"`

	AcctSvcr *BranchAndFinancialInstitutionIdentification6 `xml:"AcctSvcr"`

	RptgPrd *ReportingPeriod5 `xml:"RptgPrd"`

	RptgSeq *SequenceRange1Choice `xml:"RptgSeq"`

	ReqdTxTp *TransactionType2 `xml:"ReqdTxTp"`

	ReqdBalTp []BalanceType13 `xml:"ReqdBalTp"`

	InnerXml string `xml:",innerxml"`
}

type SequenceRange1 struct {
	XMLName xml.Name

	FrSeq Max35Text `xml:"FrSeq"`

	ToSeq Max35Text `xml:"ToSeq"`

	InnerXml string `xml:",innerxml"`
}

type SequenceRange1Choice struct {
	XMLName xml.Name

	FrSeq *Max35Text `xml:"FrSeq"`

	ToSeq *Max35Text `xml:"ToSeq"`

	FrToSeq []SequenceRange1 `xml:"FrToSeq"`

	EQSeq []Max35Text `xml:"EQSeq"`

	NEQSeq []Max35Text `xml:"NEQSeq"`

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

type TimePeriodDetails1 struct {
	XMLName xml.Name

	FrTm ISOTime `xml:"FrTm"`

	ToTm *ISOTime `xml:"ToTm"`

	InnerXml string `xml:",innerxml"`
}

type TransactionType2 struct {
	XMLName xml.Name

	Sts EntryStatus1Choice `xml:"Sts"`

	CdtDbtInd CreditDebitCode `xml:"CdtDbtInd"`

	FlrLmt []Limit2 `xml:"FlrLmt"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType float64

type ActiveOrHistoricCurrencyCode string

type AddressType2Code string

const AddressType2CodeAddr AddressType2Code = "ADDR"

const AddressType2CodePbox AddressType2Code = "PBOX"

const AddressType2CodeHome AddressType2Code = "HOME"

const AddressType2CodeBizz AddressType2Code = "BIZZ"

const AddressType2CodeMlto AddressType2Code = "MLTO"

const AddressType2CodeDlvy AddressType2Code = "DLVY"

type AnyBICDec2014Identifier string

type BICFIDec2014Identifier string

type CountryCode string

type CreditDebitCode string

const CreditDebitCodeCrdt CreditDebitCode = "CRDT"

const CreditDebitCodeDbit CreditDebitCode = "DBIT"

type Exact4AlphaNumericText string

type ExternalAccountIdentification1Code string

type ExternalBalanceSubType1Code string

type ExternalBalanceType1Code string

type ExternalCashAccountType1Code string

type ExternalClearingSystemIdentification1Code string

type ExternalEntryStatus1Code string

type ExternalFinancialInstitutionIdentification1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type ExternalProxyAccountType1Code string

type FloorLimitType1Code string

const FloorLimitType1CodeCred FloorLimitType1Code = "CRED"

const FloorLimitType1CodeDebt FloorLimitType1Code = "DEBT"

const FloorLimitType1CodeBoth FloorLimitType1Code = "BOTH"

type IBAN2007Identifier string

type ISODate string

type ISODateTime string

type ISOTime string

type LEIIdentifier string

type Max128Text string

type Max140Text string

type Max16Text string

type Max2048Text string

type Max34Text string

type Max350Text string

type Max35Text string

type Max4Text string

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

type QueryType3Code string

const QueryType3CodeAlll QueryType3Code = "ALLL"

const QueryType3CodeChng QueryType3Code = "CHNG"

const QueryType3CodeModf QueryType3Code = "MODF"