// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.004.001.11
package pacs_004_001_11

import (
	"encoding/xml"
)

// Element
type Document struct {
	XMLName xml.Name `xml:"Document"`

	PmtRtr PaymentReturnV11 `xml:",any"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	XMLName xml.Name

	IBAN *IBAN2007Identifier `xml:"IBAN"`

	Othr *GenericAccountIdentification1 `xml:"Othr"`

	InnerXml string `xml:",innerxml"`
}

type AccountSchemeName1Choice struct {
	XMLName xml.Name

	Cd *ExternalAccountIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ActiveCurrencyAndAmount struct {
	XMLName xml.Name

	Ccy ActiveCurrencyCode `xml:"Ccy,attr"`

	Text     string `xml:",chardata"`
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

type AmendmentInformationDetails14 struct {
	XMLName xml.Name

	OrgnlMndtId *Max35Text `xml:"OrgnlMndtId"`

	OrgnlCdtrSchmeId *PartyIdentification135 `xml:"OrgnlCdtrSchmeId"`

	OrgnlCdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt"`

	OrgnlCdtrAgtAcct *CashAccount40 `xml:"OrgnlCdtrAgtAcct"`

	OrgnlDbtr *PartyIdentification135 `xml:"OrgnlDbtr"`

	OrgnlDbtrAcct *CashAccount40 `xml:"OrgnlDbtrAcct"`

	OrgnlDbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt"`

	OrgnlDbtrAgtAcct *CashAccount40 `xml:"OrgnlDbtrAgtAcct"`

	OrgnlFnlColltnDt *ISODate `xml:"OrgnlFnlColltnDt"`

	OrgnlFrqcy *Frequency36Choice `xml:"OrgnlFrqcy"`

	OrgnlRsn *MandateSetupReason1Choice `xml:"OrgnlRsn"`

	OrgnlTrckgDays *Exact2NumericText `xml:"OrgnlTrckgDays"`

	InnerXml string `xml:",innerxml"`
}

type AmountType4Choice struct {
	XMLName xml.Name

	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	EqvtAmt *EquivalentAmount2 `xml:"EqvtAmt"`

	InnerXml string `xml:",innerxml"`
}

type Authorisation1Choice struct {
	XMLName xml.Name

	Cd *Authorisation1Code `xml:"Cd"`

	Prtry *Max128Text `xml:"Prtry"`

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

type CategoryPurpose1Choice struct {
	XMLName xml.Name

	Cd *ExternalCategoryPurpose1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type Charges7 struct {
	XMLName xml.Name

	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	Agt BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`

	InnerXml string `xml:",innerxml"`
}

type ClearingSystemIdentification2Choice struct {
	XMLName xml.Name

	Cd *ExternalClearingSystemIdentification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ClearingSystemIdentification3Choice struct {
	XMLName xml.Name

	Cd *ExternalCashClearingSystem1Code `xml:"Cd"`

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

type CreditTransferMandateData1 struct {
	XMLName xml.Name

	MndtId *Max35Text `xml:"MndtId"`

	Tp *MandateTypeInformation2 `xml:"Tp"`

	DtOfSgntr *ISODate `xml:"DtOfSgntr"`

	DtOfVrfctn *ISODateTime `xml:"DtOfVrfctn"`

	ElctrncSgntr *Max10KBinary `xml:"ElctrncSgntr"`

	FrstPmtDt *ISODate `xml:"FrstPmtDt"`

	FnlPmtDt *ISODate `xml:"FnlPmtDt"`

	Frqcy *Frequency36Choice `xml:"Frqcy"`

	Rsn *MandateSetupReason1Choice `xml:"Rsn"`

	InnerXml string `xml:",innerxml"`
}

type CreditTransferTransaction52 struct {
	XMLName xml.Name

	UltmtDbtr *PartyIdentification135 `xml:"UltmtDbtr"`

	InitgPty *PartyIdentification135 `xml:"InitgPty"`

	Dbtr PartyIdentification135 `xml:"Dbtr"`

	DbtrAcct *CashAccount40 `xml:"DbtrAcct"`

	DbtrAgt BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`

	DbtrAgtAcct *CashAccount40 `xml:"DbtrAgtAcct"`

	PrvsInstgAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1"`

	PrvsInstgAgt1Acct *CashAccount40 `xml:"PrvsInstgAgt1Acct"`

	PrvsInstgAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2"`

	PrvsInstgAgt2Acct *CashAccount40 `xml:"PrvsInstgAgt2Acct"`

	PrvsInstgAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3"`

	PrvsInstgAgt3Acct *CashAccount40 `xml:"PrvsInstgAgt3Acct"`

	IntrmyAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1"`

	IntrmyAgt1Acct *CashAccount40 `xml:"IntrmyAgt1Acct"`

	IntrmyAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2"`

	IntrmyAgt2Acct *CashAccount40 `xml:"IntrmyAgt2Acct"`

	IntrmyAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3"`

	IntrmyAgt3Acct *CashAccount40 `xml:"IntrmyAgt3Acct"`

	CdtrAgt BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`

	CdtrAgtAcct *CashAccount40 `xml:"CdtrAgtAcct"`

	Cdtr PartyIdentification135 `xml:"Cdtr"`

	CdtrAcct *CashAccount40 `xml:"CdtrAcct"`

	UltmtCdtr *PartyIdentification135 `xml:"UltmtCdtr"`

	InstrForCdtrAgt []InstructionForCreditorAgent3 `xml:"InstrForCdtrAgt"`

	InstrForNxtAgt []InstructionForNextAgent1 `xml:"InstrForNxtAgt"`

	Tax *TaxInformation10 `xml:"Tax"`

	RmtInf *RemittanceInformation21 `xml:"RmtInf"`

	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	InnerXml string `xml:",innerxml"`
}

type CreditorReferenceInformation2 struct {
	XMLName xml.Name

	Tp *CreditorReferenceType2 `xml:"Tp"`

	Ref *Max35Text `xml:"Ref"`

	InnerXml string `xml:",innerxml"`
}

type CreditorReferenceType1Choice struct {
	XMLName xml.Name

	Cd *DocumentType3Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type CreditorReferenceType2 struct {
	XMLName xml.Name

	CdOrPrtry CreditorReferenceType1Choice `xml:"CdOrPrtry"`

	Issr *Max35Text `xml:"Issr"`

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

type DatePeriod2 struct {
	XMLName xml.Name

	FrDt ISODate `xml:"FrDt"`

	ToDt ISODate `xml:"ToDt"`

	InnerXml string `xml:",innerxml"`
}

type DiscountAmountAndType1 struct {
	XMLName xml.Name

	Tp *DiscountAmountType1Choice `xml:"Tp"`

	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	InnerXml string `xml:",innerxml"`
}

type DiscountAmountType1Choice struct {
	XMLName xml.Name

	Cd *ExternalDiscountAmountType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type DocumentAdjustment1 struct {
	XMLName xml.Name

	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	CdtDbtInd *CreditDebitCode `xml:"CdtDbtInd"`

	Rsn *Max4Text `xml:"Rsn"`

	AddtlInf *Max140Text `xml:"AddtlInf"`

	InnerXml string `xml:",innerxml"`
}

type DocumentLineIdentification1 struct {
	XMLName xml.Name

	Tp *DocumentLineType1 `xml:"Tp"`

	Nb *Max35Text `xml:"Nb"`

	RltdDt *ISODate `xml:"RltdDt"`

	InnerXml string `xml:",innerxml"`
}

type DocumentLineInformation1 struct {
	XMLName xml.Name

	Id []DocumentLineIdentification1 `xml:"Id"`

	Desc *Max2048Text `xml:"Desc"`

	Amt *RemittanceAmount3 `xml:"Amt"`

	InnerXml string `xml:",innerxml"`
}

type DocumentLineType1 struct {
	XMLName xml.Name

	CdOrPrtry DocumentLineType1Choice `xml:"CdOrPrtry"`

	Issr *Max35Text `xml:"Issr"`

	InnerXml string `xml:",innerxml"`
}

type DocumentLineType1Choice struct {
	XMLName xml.Name

	Cd *ExternalDocumentLineType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type EquivalentAmount2 struct {
	XMLName xml.Name

	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	CcyOfTrf ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf"`

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

type Frequency36Choice struct {
	XMLName xml.Name

	Tp *Frequency6Code `xml:"Tp"`

	Prd *FrequencyPeriod1 `xml:"Prd"`

	PtInTm *FrequencyAndMoment1 `xml:"PtInTm"`

	InnerXml string `xml:",innerxml"`
}

type FrequencyAndMoment1 struct {
	XMLName xml.Name

	Tp Frequency6Code `xml:"Tp"`

	PtInTm Exact2NumericText `xml:"PtInTm"`

	InnerXml string `xml:",innerxml"`
}

type FrequencyPeriod1 struct {
	XMLName xml.Name

	Tp Frequency6Code `xml:"Tp"`

	CntPerPrd DecimalNumber `xml:"CntPerPrd"`

	InnerXml string `xml:",innerxml"`
}

type Garnishment3 struct {
	XMLName xml.Name

	Tp GarnishmentType1 `xml:"Tp"`

	Grnshee *PartyIdentification135 `xml:"Grnshee"`

	GrnshmtAdmstr *PartyIdentification135 `xml:"GrnshmtAdmstr"`

	RefNb *Max140Text `xml:"RefNb"`

	Dt *ISODate `xml:"Dt"`

	RmtdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt"`

	FmlyMdclInsrncInd *TrueFalseIndicator `xml:"FmlyMdclInsrncInd"`

	MplyeeTermntnInd *TrueFalseIndicator `xml:"MplyeeTermntnInd"`

	InnerXml string `xml:",innerxml"`
}

type GarnishmentType1 struct {
	XMLName xml.Name

	CdOrPrtry GarnishmentType1Choice `xml:"CdOrPrtry"`

	Issr *Max35Text `xml:"Issr"`

	InnerXml string `xml:",innerxml"`
}

type GarnishmentType1Choice struct {
	XMLName xml.Name

	Cd *ExternalGarnishmentType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

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

type GroupHeader99 struct {
	XMLName xml.Name

	MsgId Max35Text `xml:"MsgId"`

	CreDtTm ISODateTime `xml:"CreDtTm"`

	Authstn []Authorisation1Choice `xml:"Authstn"`

	BtchBookg *BatchBookingIndicator `xml:"BtchBookg"`

	NbOfTxs Max15NumericText `xml:"NbOfTxs"`

	CtrlSum *DecimalNumber `xml:"CtrlSum"`

	GrpRtr *TrueFalseIndicator `xml:"GrpRtr"`

	TtlRtrdIntrBkSttlmAmt *ActiveCurrencyAndAmount `xml:"TtlRtrdIntrBkSttlmAmt"`

	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt"`

	SttlmInf SettlementInstruction11 `xml:"SttlmInf"`

	PmtTpInf *PaymentTypeInformation28 `xml:"PmtTpInf"`

	InstgAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`

	InstdAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`

	InnerXml string `xml:",innerxml"`
}

type InstructionForCreditorAgent3 struct {
	XMLName xml.Name

	Cd *ExternalCreditorAgentInstruction1Code `xml:"Cd"`

	InstrInf *Max140Text `xml:"InstrInf"`

	InnerXml string `xml:",innerxml"`
}

type InstructionForNextAgent1 struct {
	XMLName xml.Name

	Cd *Instruction4Code `xml:"Cd"`

	InstrInf *Max140Text `xml:"InstrInf"`

	InnerXml string `xml:",innerxml"`
}

type LocalInstrument2Choice struct {
	XMLName xml.Name

	Cd *ExternalLocalInstrument1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type MandateClassification1Choice struct {
	XMLName xml.Name

	Cd *MandateClassification1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type MandateRelatedData2Choice struct {
	XMLName xml.Name

	DrctDbtMndt *MandateRelatedInformation15 `xml:"DrctDbtMndt"`

	CdtTrfMndt *CreditTransferMandateData1 `xml:"CdtTrfMndt"`

	InnerXml string `xml:",innerxml"`
}

type MandateRelatedInformation15 struct {
	XMLName xml.Name

	MndtId *Max35Text `xml:"MndtId"`

	DtOfSgntr *ISODate `xml:"DtOfSgntr"`

	AmdmntInd *TrueFalseIndicator `xml:"AmdmntInd"`

	AmdmntInfDtls *AmendmentInformationDetails14 `xml:"AmdmntInfDtls"`

	ElctrncSgntr *Max1025Text `xml:"ElctrncSgntr"`

	FrstColltnDt *ISODate `xml:"FrstColltnDt"`

	FnlColltnDt *ISODate `xml:"FnlColltnDt"`

	Frqcy *Frequency36Choice `xml:"Frqcy"`

	Rsn *MandateSetupReason1Choice `xml:"Rsn"`

	TrckgDays *Exact2NumericText `xml:"TrckgDays"`

	InnerXml string `xml:",innerxml"`
}

type MandateSetupReason1Choice struct {
	XMLName xml.Name

	Cd *ExternalMandateSetupReason1Code `xml:"Cd"`

	Prtry *Max70Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type MandateTypeInformation2 struct {
	XMLName xml.Name

	SvcLvl *ServiceLevel8Choice `xml:"SvcLvl"`

	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm"`

	CtgyPurp *CategoryPurpose1Choice `xml:"CtgyPurp"`

	Clssfctn *MandateClassification1Choice `xml:"Clssfctn"`

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

type OriginalGroupHeader18 struct {
	XMLName xml.Name

	OrgnlMsgId Max35Text `xml:"OrgnlMsgId"`

	OrgnlMsgNmId Max35Text `xml:"OrgnlMsgNmId"`

	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm"`

	RtrRsnInf []PaymentReturnReason6 `xml:"RtrRsnInf"`

	InnerXml string `xml:",innerxml"`
}

type OriginalGroupInformation29 struct {
	XMLName xml.Name

	OrgnlMsgId Max35Text `xml:"OrgnlMsgId"`

	OrgnlMsgNmId Max35Text `xml:"OrgnlMsgNmId"`

	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm"`

	InnerXml string `xml:",innerxml"`
}

type OriginalTransactionReference36 struct {
	XMLName xml.Name

	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	Amt *AmountType4Choice `xml:"Amt"`

	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt"`

	ReqdColltnDt *ISODate `xml:"ReqdColltnDt"`

	ReqdExctnDt *DateAndDateTime2Choice `xml:"ReqdExctnDt"`

	CdtrSchmeId *PartyIdentification135 `xml:"CdtrSchmeId"`

	SttlmInf *SettlementInstruction11 `xml:"SttlmInf"`

	PmtTpInf *PaymentTypeInformation27 `xml:"PmtTpInf"`

	PmtMtd *PaymentMethod4Code `xml:"PmtMtd"`

	MndtRltdInf *MandateRelatedData2Choice `xml:"MndtRltdInf"`

	RmtInf *RemittanceInformation21 `xml:"RmtInf"`

	UltmtDbtr *Party40Choice `xml:"UltmtDbtr"`

	Dbtr *Party40Choice `xml:"Dbtr"`

	DbtrAcct *CashAccount40 `xml:"DbtrAcct"`

	DbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`

	DbtrAgtAcct *CashAccount40 `xml:"DbtrAgtAcct"`

	CdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`

	CdtrAgtAcct *CashAccount40 `xml:"CdtrAgtAcct"`

	Cdtr *Party40Choice `xml:"Cdtr"`

	CdtrAcct *CashAccount40 `xml:"CdtrAcct"`

	UltmtCdtr *Party40Choice `xml:"UltmtCdtr"`

	Purp *Purpose2Choice `xml:"Purp"`

	UndrlygCstmrCdtTrf *CreditTransferTransaction52 `xml:"UndrlygCstmrCdtTrf"`

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

type PaymentReturnReason6 struct {
	XMLName xml.Name

	Orgtr *PartyIdentification135 `xml:"Orgtr"`

	Rsn *ReturnReason5Choice `xml:"Rsn"`

	AddtlInf []Max105Text `xml:"AddtlInf"`

	InnerXml string `xml:",innerxml"`
}

type PaymentReturnV11 struct {
	XMLName xml.Name

	GrpHdr GroupHeader99 `xml:"GrpHdr"`

	OrgnlGrpInf *OriginalGroupHeader18 `xml:"OrgnlGrpInf"`

	TxInf []PaymentTransaction133 `xml:"TxInf"`

	SplmtryData []SupplementaryData1 `xml:"SplmtryData"`

	InnerXml string `xml:",innerxml"`
}

type PaymentTransaction133 struct {
	XMLName xml.Name

	RtrId *Max35Text `xml:"RtrId"`

	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf"`

	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId"`

	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId"`

	OrgnlTxId *Max35Text `xml:"OrgnlTxId"`

	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR"`

	OrgnlClrSysRef *Max35Text `xml:"OrgnlClrSysRef"`

	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt"`

	OrgnlIntrBkSttlmDt *ISODate `xml:"OrgnlIntrBkSttlmDt"`

	PmtTpInf *PaymentTypeInformation28 `xml:"PmtTpInf"`

	RtrdIntrBkSttlmAmt ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt"`

	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt"`

	SttlmPrty *Priority3Code `xml:"SttlmPrty"`

	SttlmTmIndctn *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn"`

	SttlmTmReq *SettlementTimeRequest2 `xml:"SttlmTmReq"`

	RtrdInstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt"`

	XchgRate *BaseOneRate `xml:"XchgRate"`

	CompstnAmt *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt"`

	ChrgBr *ChargeBearerType1Code `xml:"ChrgBr"`

	ChrgsInf []Charges7 `xml:"ChrgsInf"`

	ClrSysRef *Max35Text `xml:"ClrSysRef"`

	InstgAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`

	InstdAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`

	RtrChain *TransactionParties10 `xml:"RtrChain"`

	RtrRsnInf []PaymentReturnReason6 `xml:"RtrRsnInf"`

	OrgnlTxRef *OriginalTransactionReference36 `xml:"OrgnlTxRef"`

	SplmtryData []SupplementaryData1 `xml:"SplmtryData"`

	InnerXml string `xml:",innerxml"`
}

type PaymentTypeInformation27 struct {
	XMLName xml.Name

	InstrPrty *Priority2Code `xml:"InstrPrty"`

	ClrChanl *ClearingChannel2Code `xml:"ClrChanl"`

	SvcLvl []ServiceLevel8Choice `xml:"SvcLvl"`

	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm"`

	SeqTp *SequenceType3Code `xml:"SeqTp"`

	CtgyPurp *CategoryPurpose1Choice `xml:"CtgyPurp"`

	InnerXml string `xml:",innerxml"`
}

type PaymentTypeInformation28 struct {
	XMLName xml.Name

	InstrPrty *Priority2Code `xml:"InstrPrty"`

	ClrChanl *ClearingChannel2Code `xml:"ClrChanl"`

	SvcLvl []ServiceLevel8Choice `xml:"SvcLvl"`

	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm"`

	CtgyPurp *CategoryPurpose1Choice `xml:"CtgyPurp"`

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

type Purpose2Choice struct {
	XMLName xml.Name

	Cd *ExternalPurpose1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ReferredDocumentInformation7 struct {
	XMLName xml.Name

	Tp *ReferredDocumentType4 `xml:"Tp"`

	Nb *Max35Text `xml:"Nb"`

	RltdDt *ISODate `xml:"RltdDt"`

	LineDtls []DocumentLineInformation1 `xml:"LineDtls"`

	InnerXml string `xml:",innerxml"`
}

type ReferredDocumentType3Choice struct {
	XMLName xml.Name

	Cd *DocumentType6Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ReferredDocumentType4 struct {
	XMLName xml.Name

	CdOrPrtry ReferredDocumentType3Choice `xml:"CdOrPrtry"`

	Issr *Max35Text `xml:"Issr"`

	InnerXml string `xml:",innerxml"`
}

type RemittanceAmount2 struct {
	XMLName xml.Name

	DuePyblAmt *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt"`

	DscntApldAmt []DiscountAmountAndType1 `xml:"DscntApldAmt"`

	CdtNoteAmt *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt"`

	TaxAmt []TaxAmountAndType1 `xml:"TaxAmt"`

	AdjstmntAmtAndRsn []DocumentAdjustment1 `xml:"AdjstmntAmtAndRsn"`

	RmtdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt"`

	InnerXml string `xml:",innerxml"`
}

type RemittanceAmount3 struct {
	XMLName xml.Name

	DuePyblAmt *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt"`

	DscntApldAmt []DiscountAmountAndType1 `xml:"DscntApldAmt"`

	CdtNoteAmt *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt"`

	TaxAmt []TaxAmountAndType1 `xml:"TaxAmt"`

	AdjstmntAmtAndRsn []DocumentAdjustment1 `xml:"AdjstmntAmtAndRsn"`

	RmtdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt"`

	InnerXml string `xml:",innerxml"`
}

type RemittanceInformation21 struct {
	XMLName xml.Name

	Ustrd []Max140Text `xml:"Ustrd"`

	Strd []StructuredRemittanceInformation17 `xml:"Strd"`

	InnerXml string `xml:",innerxml"`
}

type ReturnReason5Choice struct {
	XMLName xml.Name

	Cd *ExternalReturnReason1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type ServiceLevel8Choice struct {
	XMLName xml.Name

	Cd *ExternalServiceLevel1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type SettlementDateTimeIndication1 struct {
	XMLName xml.Name

	DbtDtTm *ISODateTime `xml:"DbtDtTm"`

	CdtDtTm *ISODateTime `xml:"CdtDtTm"`

	InnerXml string `xml:",innerxml"`
}

type SettlementInstruction11 struct {
	XMLName xml.Name

	SttlmMtd SettlementMethod1Code `xml:"SttlmMtd"`

	SttlmAcct *CashAccount40 `xml:"SttlmAcct"`

	ClrSys *ClearingSystemIdentification3Choice `xml:"ClrSys"`

	InstgRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt"`

	InstgRmbrsmntAgtAcct *CashAccount40 `xml:"InstgRmbrsmntAgtAcct"`

	InstdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt"`

	InstdRmbrsmntAgtAcct *CashAccount40 `xml:"InstdRmbrsmntAgtAcct"`

	ThrdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt"`

	ThrdRmbrsmntAgtAcct *CashAccount40 `xml:"ThrdRmbrsmntAgtAcct"`

	InnerXml string `xml:",innerxml"`
}

type SettlementTimeRequest2 struct {
	XMLName xml.Name

	CLSTm *ISOTime `xml:"CLSTm"`

	TillTm *ISOTime `xml:"TillTm"`

	FrTm *ISOTime `xml:"FrTm"`

	RjctTm *ISOTime `xml:"RjctTm"`

	InnerXml string `xml:",innerxml"`
}

type StructuredRemittanceInformation17 struct {
	XMLName xml.Name

	RfrdDocInf []ReferredDocumentInformation7 `xml:"RfrdDocInf"`

	RfrdDocAmt *RemittanceAmount2 `xml:"RfrdDocAmt"`

	CdtrRefInf *CreditorReferenceInformation2 `xml:"CdtrRefInf"`

	Invcr *PartyIdentification135 `xml:"Invcr"`

	Invcee *PartyIdentification135 `xml:"Invcee"`

	TaxRmt *TaxData1 `xml:"TaxRmt"`

	GrnshmtRmt *Garnishment3 `xml:"GrnshmtRmt"`

	AddtlRmtInf []Max140Text `xml:"AddtlRmtInf"`

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

type TaxAmount3 struct {
	XMLName xml.Name

	Rate *PercentageRate `xml:"Rate"`

	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt"`

	TtlAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt"`

	Dtls []TaxRecordDetails3 `xml:"Dtls"`

	InnerXml string `xml:",innerxml"`
}

type TaxAmountAndType1 struct {
	XMLName xml.Name

	Tp *TaxAmountType1Choice `xml:"Tp"`

	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	InnerXml string `xml:",innerxml"`
}

type TaxAmountType1Choice struct {
	XMLName xml.Name

	Cd *ExternalTaxAmountType1Code `xml:"Cd"`

	Prtry *Max35Text `xml:"Prtry"`

	InnerXml string `xml:",innerxml"`
}

type TaxAuthorisation1 struct {
	XMLName xml.Name

	Titl *Max35Text `xml:"Titl"`

	Nm *Max140Text `xml:"Nm"`

	InnerXml string `xml:",innerxml"`
}

type TaxData1 struct {
	XMLName xml.Name

	Cdtr *TaxParty1 `xml:"Cdtr"`

	Dbtr *TaxParty2 `xml:"Dbtr"`

	UltmtDbtr *TaxParty2 `xml:"UltmtDbtr"`

	AdmstnZone *Max35Text `xml:"AdmstnZone"`

	RefNb *Max140Text `xml:"RefNb"`

	Mtd *Max35Text `xml:"Mtd"`

	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt"`

	TtlTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt"`

	Dt *ISODate `xml:"Dt"`

	SeqNb *Number `xml:"SeqNb"`

	Rcrd []TaxRecord3 `xml:"Rcrd"`

	InnerXml string `xml:",innerxml"`
}

type TaxInformation10 struct {
	XMLName xml.Name

	Cdtr *TaxParty1 `xml:"Cdtr"`

	Dbtr *TaxParty2 `xml:"Dbtr"`

	AdmstnZone *Max35Text `xml:"AdmstnZone"`

	RefNb *Max140Text `xml:"RefNb"`

	Mtd *Max35Text `xml:"Mtd"`

	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt"`

	TtlTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt"`

	Dt *ISODate `xml:"Dt"`

	SeqNb *Number `xml:"SeqNb"`

	Rcrd []TaxRecord3 `xml:"Rcrd"`

	InnerXml string `xml:",innerxml"`
}

type TaxParty1 struct {
	XMLName xml.Name

	TaxId *Max35Text `xml:"TaxId"`

	RegnId *Max35Text `xml:"RegnId"`

	TaxTp *Max35Text `xml:"TaxTp"`

	InnerXml string `xml:",innerxml"`
}

type TaxParty2 struct {
	XMLName xml.Name

	TaxId *Max35Text `xml:"TaxId"`

	RegnId *Max35Text `xml:"RegnId"`

	TaxTp *Max35Text `xml:"TaxTp"`

	Authstn *TaxAuthorisation1 `xml:"Authstn"`

	InnerXml string `xml:",innerxml"`
}

type TaxPeriod3 struct {
	XMLName xml.Name

	Yr *ISOYear `xml:"Yr"`

	Tp *TaxRecordPeriod1Code `xml:"Tp"`

	FrToDt *DatePeriod2 `xml:"FrToDt"`

	InnerXml string `xml:",innerxml"`
}

type TaxRecord3 struct {
	XMLName xml.Name

	Tp *Max35Text `xml:"Tp"`

	Ctgy *Max35Text `xml:"Ctgy"`

	CtgyDtls *Max35Text `xml:"CtgyDtls"`

	DbtrSts *Max35Text `xml:"DbtrSts"`

	CertId *Max35Text `xml:"CertId"`

	FrmsCd *Max35Text `xml:"FrmsCd"`

	Prd *TaxPeriod3 `xml:"Prd"`

	TaxAmt *TaxAmount3 `xml:"TaxAmt"`

	AddtlInf *Max140Text `xml:"AddtlInf"`

	InnerXml string `xml:",innerxml"`
}

type TaxRecordDetails3 struct {
	XMLName xml.Name

	Prd *TaxPeriod3 `xml:"Prd"`

	Amt ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	InnerXml string `xml:",innerxml"`
}

type TransactionParties10 struct {
	XMLName xml.Name

	UltmtDbtr *Party40Choice `xml:"UltmtDbtr"`

	Dbtr Party40Choice `xml:"Dbtr"`

	DbtrAcct *CashAccount40 `xml:"DbtrAcct"`

	InitgPty *Party40Choice `xml:"InitgPty"`

	DbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`

	DbtrAgtAcct *CashAccount40 `xml:"DbtrAgtAcct"`

	PrvsInstgAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1"`

	PrvsInstgAgt1Acct *CashAccount40 `xml:"PrvsInstgAgt1Acct"`

	PrvsInstgAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2"`

	PrvsInstgAgt2Acct *CashAccount40 `xml:"PrvsInstgAgt2Acct"`

	PrvsInstgAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3"`

	PrvsInstgAgt3Acct *CashAccount40 `xml:"PrvsInstgAgt3Acct"`

	IntrmyAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1"`

	IntrmyAgt1Acct *CashAccount40 `xml:"IntrmyAgt1Acct"`

	IntrmyAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2"`

	IntrmyAgt2Acct *CashAccount40 `xml:"IntrmyAgt2Acct"`

	IntrmyAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3"`

	IntrmyAgt3Acct *CashAccount40 `xml:"IntrmyAgt3Acct"`

	CdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`

	CdtrAgtAcct *CashAccount40 `xml:"CdtrAgtAcct"`

	Cdtr Party40Choice `xml:"Cdtr"`

	CdtrAcct *CashAccount40 `xml:"CdtrAcct"`

	UltmtCdtr *Party40Choice `xml:"UltmtCdtr"`

	InnerXml string `xml:",innerxml"`
}

// XSD SimpleType declarations

type ActiveCurrencyAndAmountSimpleType float64

type ActiveCurrencyCode string

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

type Authorisation1Code string

const Authorisation1CodeAuth Authorisation1Code = "AUTH"

const Authorisation1CodeFdet Authorisation1Code = "FDET"

const Authorisation1CodeFsum Authorisation1Code = "FSUM"

const Authorisation1CodeIlev Authorisation1Code = "ILEV"

type BICFIDec2014Identifier string

type BaseOneRate float64

type BatchBookingIndicator bool

type ChargeBearerType1Code string

const ChargeBearerType1CodeDebt ChargeBearerType1Code = "DEBT"

const ChargeBearerType1CodeCred ChargeBearerType1Code = "CRED"

const ChargeBearerType1CodeShar ChargeBearerType1Code = "SHAR"

const ChargeBearerType1CodeSlev ChargeBearerType1Code = "SLEV"

type ClearingChannel2Code string

const ClearingChannel2CodeRtgs ClearingChannel2Code = "RTGS"

const ClearingChannel2CodeRtns ClearingChannel2Code = "RTNS"

const ClearingChannel2CodeMpns ClearingChannel2Code = "MPNS"

const ClearingChannel2CodeBook ClearingChannel2Code = "BOOK"

type CountryCode string

type CreditDebitCode string

const CreditDebitCodeCrdt CreditDebitCode = "CRDT"

const CreditDebitCodeDbit CreditDebitCode = "DBIT"

type DecimalNumber float64

type DocumentType3Code string

const DocumentType3CodeRadm DocumentType3Code = "RADM"

const DocumentType3CodeRpin DocumentType3Code = "RPIN"

const DocumentType3CodeFxdr DocumentType3Code = "FXDR"

const DocumentType3CodeDisp DocumentType3Code = "DISP"

const DocumentType3CodePuor DocumentType3Code = "PUOR"

const DocumentType3CodeScor DocumentType3Code = "SCOR"

type DocumentType6Code string

const DocumentType6CodeMsin DocumentType6Code = "MSIN"

const DocumentType6CodeCnfa DocumentType6Code = "CNFA"

const DocumentType6CodeDnfa DocumentType6Code = "DNFA"

const DocumentType6CodeCinv DocumentType6Code = "CINV"

const DocumentType6CodeCren DocumentType6Code = "CREN"

const DocumentType6CodeDebn DocumentType6Code = "DEBN"

const DocumentType6CodeHiri DocumentType6Code = "HIRI"

const DocumentType6CodeSbin DocumentType6Code = "SBIN"

const DocumentType6CodeCmcn DocumentType6Code = "CMCN"

const DocumentType6CodeSoac DocumentType6Code = "SOAC"

const DocumentType6CodeDisp DocumentType6Code = "DISP"

const DocumentType6CodeBold DocumentType6Code = "BOLD"

const DocumentType6CodeVchr DocumentType6Code = "VCHR"

const DocumentType6CodeAroi DocumentType6Code = "AROI"

const DocumentType6CodeTsut DocumentType6Code = "TSUT"

const DocumentType6CodePuor DocumentType6Code = "PUOR"

type Exact2NumericText string

type Exact4AlphaNumericText string

type ExternalAccountIdentification1Code string

type ExternalCashAccountType1Code string

type ExternalCashClearingSystem1Code string

type ExternalCategoryPurpose1Code string

type ExternalClearingSystemIdentification1Code string

type ExternalCreditorAgentInstruction1Code string

type ExternalDiscountAmountType1Code string

type ExternalDocumentLineType1Code string

type ExternalFinancialInstitutionIdentification1Code string

type ExternalGarnishmentType1Code string

type ExternalLocalInstrument1Code string

type ExternalMandateSetupReason1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type ExternalProxyAccountType1Code string

type ExternalPurpose1Code string

type ExternalReturnReason1Code string

type ExternalServiceLevel1Code string

type ExternalTaxAmountType1Code string

type Frequency6Code string

const Frequency6CodeYear Frequency6Code = "YEAR"

const Frequency6CodeMnth Frequency6Code = "MNTH"

const Frequency6CodeQurt Frequency6Code = "QURT"

const Frequency6CodeMian Frequency6Code = "MIAN"

const Frequency6CodeWeek Frequency6Code = "WEEK"

const Frequency6CodeDail Frequency6Code = "DAIL"

const Frequency6CodeAdho Frequency6Code = "ADHO"

const Frequency6CodeInda Frequency6Code = "INDA"

const Frequency6CodeFrtn Frequency6Code = "FRTN"

type IBAN2007Identifier string

type ISODate string

type ISODateTime string

type ISOTime string

type ISOYear string

type Instruction4Code string

const Instruction4CodePhoa Instruction4Code = "PHOA"

const Instruction4CodeTela Instruction4Code = "TELA"

type LEIIdentifier string

type MandateClassification1Code string

const MandateClassification1CodeFixe MandateClassification1Code = "FIXE"

const MandateClassification1CodeUsgb MandateClassification1Code = "USGB"

const MandateClassification1CodeVari MandateClassification1Code = "VARI"

type Max1025Text string

type Max105Text string

type Max10KBinary string

type Max128Text string

type Max140Text string

type Max15NumericText string

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

type Number float64

type PaymentMethod4Code string

const PaymentMethod4CodeChk PaymentMethod4Code = "CHK"

const PaymentMethod4CodeTrf PaymentMethod4Code = "TRF"

const PaymentMethod4CodeDd PaymentMethod4Code = "DD"

const PaymentMethod4CodeTra PaymentMethod4Code = "TRA"

type PercentageRate float64

type PhoneNumber string

type PreferredContactMethod1Code string

const PreferredContactMethod1CodeLett PreferredContactMethod1Code = "LETT"

const PreferredContactMethod1CodeMail PreferredContactMethod1Code = "MAIL"

const PreferredContactMethod1CodePhon PreferredContactMethod1Code = "PHON"

const PreferredContactMethod1CodeFaxx PreferredContactMethod1Code = "FAXX"

const PreferredContactMethod1CodeCell PreferredContactMethod1Code = "CELL"

type Priority2Code string

const Priority2CodeHigh Priority2Code = "HIGH"

const Priority2CodeNorm Priority2Code = "NORM"

type Priority3Code string

const Priority3CodeUrgt Priority3Code = "URGT"

const Priority3CodeHigh Priority3Code = "HIGH"

const Priority3CodeNorm Priority3Code = "NORM"

type SequenceType3Code string

const SequenceType3CodeFrst SequenceType3Code = "FRST"

const SequenceType3CodeRcur SequenceType3Code = "RCUR"

const SequenceType3CodeFnal SequenceType3Code = "FNAL"

const SequenceType3CodeOoff SequenceType3Code = "OOFF"

const SequenceType3CodeRpre SequenceType3Code = "RPRE"

type SettlementMethod1Code string

const SettlementMethod1CodeInda SettlementMethod1Code = "INDA"

const SettlementMethod1CodeInga SettlementMethod1Code = "INGA"

const SettlementMethod1CodeCove SettlementMethod1Code = "COVE"

const SettlementMethod1CodeClrg SettlementMethod1Code = "CLRG"

type TaxRecordPeriod1Code string

const TaxRecordPeriod1CodeMm01 TaxRecordPeriod1Code = "MM01"

const TaxRecordPeriod1CodeMm02 TaxRecordPeriod1Code = "MM02"

const TaxRecordPeriod1CodeMm03 TaxRecordPeriod1Code = "MM03"

const TaxRecordPeriod1CodeMm04 TaxRecordPeriod1Code = "MM04"

const TaxRecordPeriod1CodeMm05 TaxRecordPeriod1Code = "MM05"

const TaxRecordPeriod1CodeMm06 TaxRecordPeriod1Code = "MM06"

const TaxRecordPeriod1CodeMm07 TaxRecordPeriod1Code = "MM07"

const TaxRecordPeriod1CodeMm08 TaxRecordPeriod1Code = "MM08"

const TaxRecordPeriod1CodeMm09 TaxRecordPeriod1Code = "MM09"

const TaxRecordPeriod1CodeMm10 TaxRecordPeriod1Code = "MM10"

const TaxRecordPeriod1CodeMm11 TaxRecordPeriod1Code = "MM11"

const TaxRecordPeriod1CodeMm12 TaxRecordPeriod1Code = "MM12"

const TaxRecordPeriod1CodeQtr1 TaxRecordPeriod1Code = "QTR1"

const TaxRecordPeriod1CodeQtr2 TaxRecordPeriod1Code = "QTR2"

const TaxRecordPeriod1CodeQtr3 TaxRecordPeriod1Code = "QTR3"

const TaxRecordPeriod1CodeQtr4 TaxRecordPeriod1Code = "QTR4"

const TaxRecordPeriod1CodeHlf1 TaxRecordPeriod1Code = "HLF1"

const TaxRecordPeriod1CodeHlf2 TaxRecordPeriod1Code = "HLF2"

type TrueFalseIndicator bool

type UUIDv4Identifier string
