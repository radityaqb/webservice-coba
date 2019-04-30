package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tokopedia/campaign-engine/domain/thirdparty/galadriel"
)

func main() {
	initRoutes()
}

func initRoutes() {
	// http.HandleFunc("/get", handleGet)
	// http.HandleFunc("/get-json-1", handleGetJson1)
	// http.HandleFunc("/get-json-2", handleGetJson2)
	// http.HandleFunc("/get-json-3", handleGetJson3)
	// http.HandleFunc("/post", handlePost)
	// http.HandleFunc("/post-2", handlePost2)

	// fmt.Println("SERVING in 8181...")
	// http.ListenAndServe(":8181", nil)

	// x := "aaaaa\nbbbb"
	// ioutil.WriteFile(fmt.Sprint(time.Now().Unix())+"-fallback", []byte(x), 0644)
	// ioutil.ReadFile()
	// pattern := "*-fallback"
	// matches, err := filepath.Glob(pattern)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(matches)

	// // i love anisa
	// var (
	// 	vUses      []UseCodeInput
	// 	nSuccesses []UseCodeInput
	// )

	// type gabung struct {
	// 	vUse     UseCodeInput
	// 	nsuccess UseCodeInput
	// }

	// // VALIDATE USE
	// file, err := os.Open("/Users/nakama/Downloads/validate_use.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	// fmt.Println("MASUK SINI")

	// 	teks := scanner.Text()

	// 	teks = teks[:len(teks)-3]

	// 	var vX UseCodeInput
	// 	err := json.Unmarshal([]byte(teks), &vX)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	vUses = append(vUses, vX)
	// }

	// // NOTIFY USE SUCCESS
	// filez, err := os.Open("/Users/nakama/Downloads/notify_use_success.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer filez.Close()

	// scannerz := bufio.NewScanner(filez)
	// for scannerz.Scan() {
	// 	// fmt.Println("MASUK SINI")
	// 	teks := scannerz.Text()

	// 	teks = teks[:len(teks)-3]

	// 	var vX UseCodeInput
	// 	err := json.Unmarshal([]byte(teks), &vX)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	nSuccesses = append(nSuccesses, vX)
	// }

	// // fmt.Println(vUses)
	// // fmt.Println(nSuccesses)

	// var jancuks []gabung
	// for _, nSuccess := range nSuccesses {
	// 	// fmt.Println("MASUK SINI")
	// 	for _, vUse := range vUses {
	// 		// fmt.Println("MASUK SINIXXX")
	// 		if nSuccess.Data.UserData.UserID == vUse.Data.UserData.UserID {//&& nSuccess.Data.PaymentAmount == vUse.Data.GrandTotal {
	// 			jancuk := gabung{
	// 				vUse:     vUse,
	// 				nsuccess: nSuccess,
	// 			}
	// 			jancuks = append(jancuks, jancuk)
	// 			break
	// 		}
	// 	}
	// }

	// // fmt.Println(jancuks)

	// fmt.Println(len(jancuks))

	// for _, j := range jancuks {
	// 	x, _ := json.Marshal(j.vUse)
	// 	y, _ := json.Marshal(j.nsuccess)

	// 	fmt.Println(string(x), "~~~~~", string(y))
	// }

	// TEST
	file, err := os.Open("DOR.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Println("redis-cli -h campaign2.redis.singapore.rds.aliyuncs.com RPUSH ce:promousage '", scanner.Text(), "'")
		teks := scanner.Text()

		x := strings.Split(teks, `"~~~~~"`)
		validateUse := x[0]
		validateUse = validateUse[1:]
		notifySuccess := x[1]
		notifySuccess = notifySuccess[:len(notifySuccess)-1]

		promoCodeID, _ := hitValidateUse(validateUse)

		hitNotifySuccess(notifySuccess, promoCodeID)

		// fmt.Println(teks)
	}
}

func hitNotifySuccess(nSuccess string, promoCodeID int64) {
	// jsonReq, _ := json.Marshal([]byte(nSuccess))

	var notifSuccess UseCodeInput
	json.Unmarshal([]byte(nSuccess), &notifSuccess)

	notifSuccess.Data.PromoCodeSessionID = promoCodeID

	jsonReq, _ := json.Marshal(notifSuccess)

	fmt.Println("Request ", nSuccess)

	c := http.Client{Timeout: 1 * time.Second}

	req, _ := http.NewRequest("POST", "http://172.21.43.158/promocode/notify/use/success", bytes.NewBuffer(jsonReq))

	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response ", string(data))

	var xxxx CheckVoucherCodeResponseSamaData
	json.Unmarshal(data, &xxxx)

	return
}

func hitValidateUse(vUse string) (int64, error) {
	// jsonReq, _ := json.Marshal([]byte(vUse))

	fmt.Println("Request ", vUse)

	c := http.Client{Timeout: 1 * time.Second}

	req, _ := http.NewRequest("POST", "http://172.21.43.158/promocode/validate/use", bytes.NewBuffer([]byte(vUse)))

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Response ", string(data))

	var xxxx CheckVoucherCodeResponseSamaData
	json.Unmarshal(data, &xxxx)

	return xxxx.Data.PromoCodeSessionID, nil
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HALO DUNIA"))
}

func handleGetJson1(w http.ResponseWriter, r *http.Request) {
	ret := `{"id":1234,"name":"radit"}`
	w.Write([]byte(ret))
}

func handleGetJson2(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	res := response{
		ID:   1234,
		Name: "radit",
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(data)
}

func handleGetJson3(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	queryValues := r.URL.Query()

	idstring := queryValues.Get("id")
	id, _ := strconv.Atoi(idstring)
	name := queryValues.Get("name")

	res := response{
		ID:   id,
		Name: name,
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(data)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	type (
		request struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}

		response struct {
			Message string `json:"message"`
		}
	)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var req request
	if err = json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	msg := fmt.Sprintf("halo, nama saya %s, id saya %d", req.Name, req.ID)
	res := response{
		Message: msg,
	}

	dataresp, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(dataresp)
}

func handlePost2(w http.ResponseWriter, r *http.Request) {
	type dataStruct struct {
		ID   int    `json:"id"`
		Kata string `json:"kata"`
	}

	type (
		request struct {
			Kata string `json:"kata"`
		}

		response struct {
			Data []dataStruct `json:"data"`
		}
	)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var req request
	if err = json.Unmarshal(data, &req); err != nil {
		log.Fatal(err)
	}

	s := strings.Split(req.Kata, ",")

	res := []dataStruct{}
	for i, selem := range s {
		temp := dataStruct{
			ID:   i,
			Kata: selem,
		}
		res = append(res, temp)
	}

	var trueRes response
	trueRes.Data = res

	dataresp, err := json.Marshal(trueRes)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(dataresp)
}

// Input data for use code flow
type (
	UseCodeInput struct {
		Data UseCodeData `json:"data"`
	}

	UseCodeData struct {
		Codes              []string               `json:"codes"`
		CurrentApplyCode   UseCodeApplyInfo       `json:"current_apply_code"`
		PromoCodeSessionID int64                  `json:"promo_code_id"`
		UserData           UseCodeUserData        `json:"user_data"`
		PaymentInfo        UseCodePaymentInfo     `json:"payment_info"`
		MetaData           map[string]interface{} `json:"meta_data"` // should I use map?
		GrandTotal         float64                `json:"grand_total"`
		PaymentAmount      float64                `json:"payment_amount"`

		// service id and secret key
		ServiceID int64  `json:"service_id"`
		SecretKey string `json:"secret_key"`

		// promocode related
		IsFirstStep bool `json:"is_first_step"`
		Book        bool `json:"book"`
		IsSuggested bool `json:"is_suggested"`
		IsAppRoot   bool `json:"is_app_root"`
		SkipApply   bool `json:"skip_apply"`

		// rule engine related
		Campaign []UseCodeCampaignData `json:"campaign_data"`

		// after payment
		PaymentID                int64 `json:"payment_id"`
		FinishedOrderID          int64 `json:"finished_order_id"`
		IsSkipCheckNotifySuccess bool  `json:"is_skip_check_notify_success"`

		IsSkipFraudCheck bool      `json:"is_skip_fraud_check"`
		FraudInfo        FraudInfo `json:"fraud_info"`

		// others
		Language string `json:"language"`
		State    string `json:"state"`
	}

	UseCodeUserData struct {
		UserID                int64                  `json:"user_id"`
		Name                  string                 `json:"name"`
		Email                 string                 `json:"email"`
		Msisdn                string                 `json:"msisdn"`
		MsisdnVerified        bool                   `json:"msisdn_verified"`
		IsQcAccount           bool                   `json:"is_qc_acc"`
		AppVersion            string                 `json:"app_version"`
		UserAgent             string                 `json:"user_agent"`
		IPAddress             string                 `json:"ip_address"`
		AdsID                 string                 `json:"advertisement_id"`
		DeviceType            string                 `json:"device_type"`
		DeviceID              string                 `json:"device_id"`
		AddressDetail         map[string]interface{} `json:"address_detail"`
		UserTransactionDetail map[string]interface{} `json:"user_transaction_detail"`
	}

	UseCodePaymentInfo struct {
		ScroogeGatewayID      int64   `json:"scrooge_gateway_id"`
		ScroogeGatewayCode    string  `json:"scrooge_gateway_code"`
		ScroogeGatewayValue   string  `json:"scrooge_gateway_value"`
		CreditCardNumber      string  `json:"credit_card_number"`
		CreditCardExpiryMonth int     `json:"cc_exp_month"`
		CreditCardExpiryYear  int     `json:"cc_exp_year"`
		CreditCardHash        string  `json:"cc_hash"`
		OvoCashAmount         float64 `json:"ovo_cash_amount"`
		OvoPointsAmount       float64 `json:"ovo_points_amount"`
	}

	UseCodeCampaignData struct {
		Code        string                 `json:"code"`
		PromoCodeID int64                  `json:"promo_code_id"`
		PromoID     int64                  `json:"promo_id"`
		RuleIDs     []int64                `json:"rule_ids"`
		DoGaladriel []string               `json:"do_galadriel"`
		FlowState   int                    `json:"flow_state"`
		MetaData    map[string]interface{} `json:"meta_data"`
	}

	PromoStackAbuse struct {
		PromoID       int64 `json:"promo_id"`
		IsPromoAbuser int   `json:"is_promo_abuser"`
	}

	UseCodeServiceData struct {
		CategoryCode int    `json:"category_code"`
		ProductCode  string `json:"product_code"`
	}

	UseCodeApplyInfo struct {
		Code string `json:"code"`
		Type string `json:"type"`
	}
)

// Validate Use Response
type (

	// struct that returned to clients
	CheckVoucherCodeResponseSamaData struct {
		Data CheckVoucherCodeResponse `json:"data"`
	}

	CheckVoucherCodeResponse struct {
		GlobalSuccess                  bool                  `json:"global_success"`
		Success                        bool                  `json:"success"`
		Message                        UseCodeMessage        `json:"message"`
		Codes                          []string              `json:"codes"`
		PromoCodeSessionID             int64                 `json:"promo_code_id"`
		TitleDescription               string                `json:"title_description"`
		DiscountAmount                 float64               `json:"discount_amount"`
		CashbackTokoCashAmount         float64               `json:"cashback_wallet_amount"`
		CashbackMultiplier             int                   `json:"cashback_multiplier"`
		CashbackAdvocateReferralAmount float64               `json:"cashback_advocate_referral_amount"`
		InvoiceDescription             string                `json:"invoice_description"`
		IsCoupon                       int                   `json:"is_coupon"`
		GatewayID                      string                `json:"gateway_id"`
		ClashingInfoDetail             ClashingInfoDetail    `json:"clashing_info_detail,omitempty"`
		BenefitDetails                 []BenefitOrderDetail  `json:"benefit_details"`
		VoucherOrders                  []UseCodeVoucherOrder `json:"voucher_orders"`
		BenefitSummaryInfo             BenefitSummaryInfo    `json:"benefit_summary_info"`

		V Validate
	}

	BenefitOrderDetail struct {
		Code                  string                 `json:"code"`
		OrderID               int64                  `json:"order_id"`
		UniqueID              string                 `json:"unique_id"`
		DiscountAmount        float64                `json:"discount_amount"`
		DiscountDataType      string                 `json:"discount_data_type"`
		CashbackAmount        float64                `json:"cashback_amount"`
		BenefitProductDetails []BenefitProductDetail `json:"benefit_product_details"`
	}

	BenefitProductDetail struct {
		ProductID      int64                       `json:"product_id"`
		CashbackAmount float64                     `json:"cashback_amount"`
		DiscountAmount float64                     `json:"discount_amount"`
		Promo          []BenefitProductDetailPromo `json:"promo"`
	}

	BenefitProductDetailPromo struct {
		Code           string `json:"code"`
		Tier           string `json:"tier"`
		TieringItem    string
		DiscountAmount float64 `json:"discount_amount"`
		CashbackAmount float64 `json:"cashback_amount"`
		ShowTag        bool    `json:"show_tag"`
	}

	BenefitSummaryInfo struct {
		FinalBenefitText      string           `json:"final_benefit_text"`
		FinalBenefitAmountStr string           `json:"final_benefit_amount_str"`
		FinalBenefitAmount    float64          `json:"final_benefit_amount"`
		Summaries             []BenefitSummary `json:"summaries"`
	}

	BenefitSummary struct {
		Description string  `json:"description"`
		Type        string  `json:"type"`
		AmountStr   string  `json:"amount_str"`
		Amount      float64 `json:"amount"`
	}

	UseCodeVoucherOrder struct {
		Code                   string               `json:"code"`
		Success                bool                 `json:"success"`
		UniqueID               string               `json:"unique_id"`
		OrderID                int64                `json:"order_id"`
		ShopID                 int64                `json:"shop_id"`
		IsPO                   bool                 `json:"is_po"`
		Duration               string               `json:"duration"`
		WarehouseID            int64                `json:"warehouse_id"`
		AddressID              int64                `json:"address_id"`
		Type                   string               `json:"type"`
		CashbackTokoCashAmount float64              `json:"cashback_wallet_amount"`
		CashbackMultiplier     int                  `json:"cashback_multiplier"`
		DiscountAmount         float64              `json:"discount_amount"`
		TitleDescription       string               `json:"title_description"`
		InvoiceDescription     string               `json:"invoice_description"`
		Message                UseCodeMessage       `json:"message"`
		BenefitDetails         []BenefitOrderDetail `json:"benefit_details"`

		// struct validate
		V Validate
	}

	// clash
	ClashingInfoDetail struct {
		IsClashedPromos bool          `json:"is_clashed_promos"`
		ClashReason     string        `json:"clash_reason"`
		ClashMessage    string        `json:"clash_message"`
		Options         []ClashOption `json:"options"`
	}
	ClashOption struct {
		VoucherOrders []ClashVoucherOrder `json:"voucher_orders"`
	}
	ClashVoucherOrder struct {
		Code             string  `json:"code"`
		UniqueID         string  `json:"unique_id"`
		PromoName        string  `json:"promo_name"`
		PotentialBenefit float64 `json:"potential_benefit"`
	}
)

// Validate Payment
type (
	ValidatePaymentRequest struct {
		Data ValidatePaymentData `json:"data"`
	}

	ValidatePaymentData struct {
		PromoCodeID         int64              `json:"promo_code_id"`
		PaymentID           int64              `json:"payment_id"`
		UserData            UseCodeUserData    `json:"user_data"`
		PaymentInfo         UseCodePaymentInfo `json:"payment_info"`
		ScroogeMerchantCode string             `json:"scrooge_merchant_code"`
		ServiceData         UseCodeServiceData `json:"service_data"`
		Language            string             `json:"language"`
		PaymentAmount       float64            `json:"payment_amount"`
		ServiceID           int                `json:"service_id"`
	}

	ValidatePaymentResponse struct {
		Success     bool   `json:"success"`
		PromoCodeID int64  `json:"promo_code_id"`
		Error       string `json:"error"`
	}
)

// Validate Unique ID Response
type (
	ValidateUniqueIDResponse struct {
		Success bool `json:"success"`
	}
)

// Notify Use Response
type (
	NotifyUseSuccessResponse struct {
		Success        bool                  `json:"success"`
		PromoCodeID    int64                 `json:"promo_code_id"`
		Message        string                `json:"message"`
		BenefitDetails []BenefitOrderDetail  `json:"benefit_details"`
		VourcherOrders []UseCodeVoucherOrder `json:"voucher_orders"`
	}

	NotifyUseSuccessOrderDetail struct {
		OrderID             int64                          `json:"order_id"`
		PromoType           int64                          `json:"promo_type"`
		TotalCashbackWallet float64                        `json:"total_cashback_wallet"`
		TotalDiscount       float64                        `json:"total_discount"`
		SummaryPromo        []NotifyUseSuccessSummaryPromo `json:"summary_promo"`
	}

	NotifyUseSuccessSummaryPromo struct {
		Name                 string  `json:"name"`
		IsCoupon             bool    `json:"is_coupon"`
		ShowCashbackAmount   bool    `json:"show_cashback_amount"`
		ShowDiscountAmount   bool    `json:"show_discount_amount"`
		CashbackWalletAmount float64 `json:"cashback_wallet_amount"`
		DiscountAmount       float64 `json:"discount_amount"`
	}

	NotifyUseSuccessAffectedProduct struct {
		ProductID int64                       `json:"product_id"`
		Promo     []NotifyUseSuccessPromoInfo `json:"promo"`
	}

	NotifyUseSuccessPromoInfo struct {
		Name                 string  `json:"name"`
		CashbackWalletAmount float64 `json:"cashback_wallet_amount"`
		DiscountAmount       float64 `json:"discount_amount"`
		ShowTag              bool    `json:"show_tag"`
	}

	UseCodeMessage struct {
		State string `json:"state"`
		Color string `json:"color"`
		Text  string `json:"text"`
	}

	AutoApplySaveRequest struct {
		Data struct {
			UserID    int64                  `json:"user_id"`
			ServiceID int64                  `json:"service_id"`
			Code      string                 `json:"code"`
			MetaData  map[string]interface{} `json:"meta_data"`
		} `json:"data"`
	}
)

// Do Refund Promo Code
type (

	// Refund Promo Code Request Client struct
	RefundPromoCodeInput struct {
		Data RefundData `json:"data"`
	}

	RefundData struct {
		RefundedOrderID int64           `json:"order_id"`
		PaymentID       int64           `json:"payment_id"`
		RefundAmount    float64         `json:"total_refund"`
		InvoiceRefNum   string          `json:"invoice_ref_num"`
		UserData        UseCodeUserData `json:"user_data"`
	}

	// Refund Promo Code Response struct
	RefundPromoCodeResponse struct {
		RefundStatus RefundStatus `json:"refund_status"`
		DetailData   []RefundCode `json:"codes"`
	}

	RefundStatus struct {
		RefundSuccess bool   `json:"refund_success"`
		Message       string `json:"message"`
	}

	RefundCode struct {
		Code           string `json:"code"`
		IsSuccess      bool   `json:"is_success"`
		Message        string `json:"message"`
		IsRefundCoupon bool   `json:"is_refund_coupon"`
	}

	// Refund request for Gala
	RefundRequest struct {
		Campaign []UseCodeCampaignData `json:"campaign"`
		UserData UseCodeUserData       `json:"user_data"`
	}
)

// Get Benefit Info
type (
	InputDataGetBenefit struct {
		Data struct {
			ServiceID       int64                  `json:"service_id"`
			PaymentID       int64                  `json:"payment_id"`
			FinishedOrderID int64                  `json:"finished_order_id"`
			PromoCodeID     int64                  `json:"promo_code_id"`
			SecretKey       string                 `json:"secret_key"`
			State           string                 `json:"state"`
			UserData        UseCodeUserData        `json:"user_data"`
			FraudInfo       FraudInfo              `json:"fraud_info"`
			MetaData        map[string]interface{} `json:"meta_data"`
		}
	}

	MetaData struct {
		Orders             []Order `json:"orders"`
		FirstShippingPrice float64 `json:"first_shipping_price"`
		TotalShippingPrice float64 `json:"total_shipping_price"`
	}

	Order struct {
		OrderID           int64                 `json:"order_id"`
		OrderStatus       int                   `json:"order_status"`
		RefundAmount      int64                 `json:"refund_amount"`
		TotalProductPrice float64               `json:"total_product_price"`
		ProductDetails    []ProductDetails      `json:"product_details"`
		Campaign          []UseCodeCampaignData `json:"campaign_data"`
		ShopID            int64                 `json:"shop_id"`
		IsGoldShop        bool                  `json:"is_gold_shop"`
		ShippingID        int64                 `json:"shipping_id"`
		SpID              int64                 `json:"sp_id"`
		ShippingPrice     float64               `json:"shipping_price"`
		InsurancePrice    float64               `json:"insurance_price"`
		AdditionalFee     float64               `json:"additional_fee"`
		DestProvince      int                   `json:"dest_province"`
		DestCity          int                   `json:"dest_city"`
		TotalWeight       float64               `json:"total_weight"`
		ReturnedAmount    float64               `json:"returned_amount"`
		Address           Address               `json:"address"`
		InvoiceRefNum     string                `json:"invoice_ref_num"`
		IsPowerBadge      bool                  `json:"is_power_badge"`
		IsOfficialStore   bool                  `json:"is_official_store"`
	}

	Address struct {
		ID           int64  `json:"id"`
		ProvinceID   int64  `json:"province"`
		CityID       int64  `json:"city"`
		DistrictID   int64  `json:"district"`
		PostalCode   string `json:"postal"`
		AddressName  string `json:"address_name"`
		Phone        string `json:"phone"`
		ReceiverName string `json:"receiver_name"`
		CountryName  string `json:"country_name"`
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
	}

	ProductDetails struct {
		EtalaseID    int64 `json:"etalase_id"`
		ProductID    int64 `json:"product_id"`
		Quantity     int   `json:"quantity"`
		CategoryID   int   `json:"category_id"`
		PricePerItem int64 `json:"price_per_item"`
		TotalPrice   int64 `json:"total_price"`
	}

	GetBenefitInfoResponse struct {
		PaymentID int64    `json:"payment_id"`
		OrderID   int64    `json:"order_id"`
		Promos    []Promoz `json:"promos"`
	}

	Promoz struct {
		Code           string           `json:"code"`
		CashbackAmount float64          `json:"cashback_amount"`
		DiscountAmount float64          `json:"discount_amount"`
		BinaryType     int64            `json:"binary_type"`
		Budget         galadriel.Budget `json:"budget"`
	}

	PcuaomOrderDetail struct {
		ProductID int64   `json:"product_id"`
		Promo     []Promo `json:"promo"`
	}
)

type FraudInfo struct {
	DropshipAsBuyer int               `json:"dropship_as_buyer"`
	IsPromoAbuser   int               `json:"is_promo_abuser"`
	IsInvalid       int               `json:"is_invalid"`
	Status          interface{}       `json:"status"`
	PromoStackAbuse []PromoStackAbuse `json:"promo_stack_abuse"`
	Source          string            `json:"source"`
}

// Get Invoice Response
type (
	GetInvoiceResponse struct {
		Data struct {
			Success     bool                    `json:"success"`
			PromoCodeID int                     `json:"promo_code_id"`
			Message     string                  `json:"message"`
			OrderDetail []GetInvoiceOrderDetail `json:"order_detail"`
		} `json:"data"`
	}

	GetInvoiceOrderDetail struct {
		OrderID         int64                       `json:"order_id"`
		PromoType       int                         `json:"promo_type"`
		TotalCashback   float64                     `json:"total_cashback"`
		TotalDiscount   float64                     `json:"total_discount"`
		SummaryPromo    []GetInvoiceSummaryPromo    `json:"summary_promo"`
		ExtraInfo       string                      `json:"extra_info"`
		AffectedProduct []GetInvoiceAffectedProduct `json:"affected_product"`
	}

	GetInvoiceSummaryPromo struct {
		Name               string  `json:"name"`
		IsCoupon           bool    `json:"is_coupon"`
		ShowCashbackAmount bool    `json:"show_cashback_amount"`
		ShowDiscountAmount bool    `json:"show_discount_amount"`
		CashbackAmount     float64 `json:"cashback_amount"`
		DiscountAmount     float64 `json:"discount_amount"`
		InvoiceDesc        string  `json:"invoice_desc"`
		PromoType          int     `json:"promo_type"`
	}

	GetInvoiceAffectedProduct struct {
		ProductID int64                 `json:"product_id"`
		Promo     []GetInvoicePromoInfo `json:"promo"`
	}

	GetInvoicePromoInfo struct {
		Name                 string  `json:"name"`
		Code                 string  `json:"code,omitempty"`
		CashbackWalletAmount float64 `json:"cashback_amount"`
		DiscountAmount       float64 `json:"discount_amount"`
		ShowTag              bool    `json:"show_tag"`
	}

	DetailPromo struct {
		Name           string `json:"name"`
		Code           string `json:"code"`
		Tier           string
		TieringItem    string
		DiscountAmount int
		CashbackAmount int
		PromoType      int
		IsCoupon       bool
	}

	GetInvoiceOrderDetailInOrder struct {
		Code      string
		ProductID int64
		Promo     struct {
			ShowTag bool `json:"show_tag"`
		} `json:"promo"`
	}

	OrderDetail struct {
		ProductDetail []AffectedProduct
	}

	AffectedProduct struct {
		ProductID    int64                 `json:"product_id"`
		ProductPrice int                   `json:"product_price"`
		Promo        []GetInvoicePromoInfo `json:"promo"`
	}
)

type (
	OrderMap struct {
		CodesData map[string]CodeVoucherOrder `json:"code_data"`
	}
	CodeVoucherOrder struct {
		DiscountAmount         float64 `json:"discount_amount"`
		CashbackAmount         float64 `json:"cashback_amount"`
		IsUsingTokopediaBudget bool    `json:"is_using_tokopedia_budget"`
		// struct validate
		V Validate
	}
)

type (
	NotifyRejectResponse struct {
		Data struct {
			Success bool `json:"success"`
		} `json:"data"`
	}

	NotifyFinishSuccessResponse struct {
		Success   bool  `json:"success"`
		PaymentID int64 `json:"payment_id"`
		OrderID   int64 `json:"order_id"`
	}
)

type (
	IrisParam struct {
		Resp         interface{}
		UseCodeInput UseCodeInput
		IsSuccess    bool
		// EventDataErr irisDmn.IrisEventData
		IrisState int
		FlowState int
		// Option       irisDmn.IrisOption
		ErrorProcess error
	}
)

//Update Coupon Parameter
type (
	CouponUpdateParam struct {
		PromoCodeID int64  `json:"promo_code_id"`
		Code        string `json:"code"`
		UserID      int64  `json:"user_id"`
		PromoID     int    `json:"promo_id"`
		Status      int    `json:"status"`
		ExpiredAt   int64  `json:"expired_at"`
	}
)

type (
	Promo struct {
		ID             int64   `json:"id"                db:"id"`
		Name           string  `json:"name"              db:"name"`
		RuleIDs        []int64 `json:"rule_ids"          db:"rules_id"`
		Public         bool    `json:"public"            db:"public"`
		Quota          int     `json:"quota"             db:"quota"`
		Status         int8    `json:"status"            db:"status"`
		BaseCode       string  `json:"base_code"         db:"base_code"`
		MessageSuccess string  `json:"message_success"   db:"message_success"`
		IsBackdoor     bool    `json:"is_backdoor"       db:"is_backdoor"`
		PromoType      int     `json:"promo_type"        db:"promo_type"`
		ApproverType   int     `json:"approver_type"     db:"approver_type"`
		// ApprovalUserIDs []int64 `json:"approval_user_ids" db:"approval_user_ids`
		PromoTitle               string `json:"promo_title"                 db:"promo_title"`
		Description              string `json:"description"                 db:"description"`
		Metadata                 string `json:"metadata"                    db:"metadata"`
		RuleEngineRegistrationID int64  `json:"rule_engine_registration_id" db:"rule_engine_registration_id"`
		// PromoTag           PromoTag        `jsonapi:"promo_tag" json:"promo_tag"`
	}
)

type (
	TimeWindow struct {
		ID        int64     `json:"id"`
		PromoID   int64     `json:"promo_id"`
		Name      string    `json:"name"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
		// IsQcAcc   int       `json:"is_qc_acc"`
	}

	TimeWindows struct {
		PromoID int64        `json:"promo_id"`
		TWs     []TimeWindow `json:"timewindow"`
	}
)

type (
	Budget struct {
		ID             int64          `json:"budget_id"        db:"budget_id"`
		PromoID        int64          `json:"promo_id"         db:"promo_id"`
		Tokopedia      int64          `json:"tokopedia"        db:"tokopedia"`
		Partner        int64          `json:"partner"          db:"partner"`
		Notes          string         `json:"notes"            db:"notes"`
		IsMultiBenefit bool           `json:"is_multi_benefit" db:"is_multi_benefit"`
		Details        []BudgetDetail `json:"budget_details"`
	}

	BudgetDetail struct {
		ID          int64   `json:"budget_details_id" db:"budget_details_id"`
		BudgetID    int64   `json:"budget_id" db:"budget_id"`
		BudgetType  int64   `json:"budget_type" db:"notes"`
		Notes       string  `json:"notes" db:"notes"`
		BrandID     []int64 `json:"brand_id" db:"brand_id"`
		Value       int64   `json:"value" db:"value"`
		PartnerID   int64   `json:"partner_id" db:"partner_id"`
		BenefitType int64   `json:"benefit_type" db:"benefit_type"`
	}
)

// promocode data includes data that required to do things
type (
	Validate struct {
		Code         string
		ReferralCode string
		Level        int
		U            Usage

		PromoCode   PromoCode
		Promo       Promo
		Budget      Budget
		TimeWindows TimeWindows

		// response from rule engine append here too
		GatewayID int64
	}

	Usage struct {
		PromoCodeUsage                   PromoCodeUsage `json:"promo_code_usage"`
		PromoCodeUsagePayment            PromoCodeUsagePayment
		PromoCodeUsageAffectedOrderMixes []PromoCodeUsageAffectedOrderMix
	}
)

type (
	PromoCodeUsage struct {
		ID             int64     `json:"id" db:"id"`
		PromoID        int64     `json:"promo_id" db:"promo_id"`
		PromoCodeID    int64     `json:"promo_code_id" db:"promo_code_id"`
		ServiceID      int64     `json:"service_id" db:"service_id"`
		Code           string    `json:"code" db:"code"`
		Status         int       `json:"status" db:"status"`
		PaymentID      int64     `json:"payment_id" db:"payment_id"`
		ConfirmationID int64     `json:"confirmation_id" db:"confirmation_id"`
		PaymentAmount  float64   `json:"payment_amount" db:"payment_amount"`
		CashbackAmount float64   `json:"cashback_amount" db:"cashback_amount"`
		DiscountAmount float64   `json:"discount_amount" db:"discount_amount"`
		ExtraAmount    float64   `json:"extra_amount" db:"extra_amount"`
		SaldoAmount    float64   `json:"saldo_amount" db:"saldo_amount"`
		WalletAmount   float64   `json:"wallet_amount" db:"wallet_amount"`
		CreateTime     time.Time `json:"create_time" db:"create_time"`
		UsageTime      time.Time `json:"usage_time" db:"usage_time"`
		UsageUser      int64     `json:"usage_user" db:"usage_user"`
		OwnerEmail     string    `json:"owner_email" db:"owner_email"`
		Msisdn         string    `json:"msisdn" db:"msisdn"`
		IpAddress      string    `json:"ip" db:"ip"`
		GatewayCode    string    `json:"gateway_code" db:"gateway_code"`
		IsQcAcc        int       `json:"is_qc_acc" db:"is_qc_acc"`
	}
)

type (
	PromoCodeUsagePayment struct {
		ID               int64                       `json:"id"`
		PromoCodeUsageID int64                       `json:"promo_code_usage_id"`
		PaymentID        int64                       `json:"payment_id"`
		PaymentDetail    PromoCodeUsagePaymentDetail `json:"promo_code_usage_payment_detail"`
	}

	PromoCodeUsagePaymentDetail struct {
		PaymentAmount   float64 `json:"payment_amount"`
		OvoCashAmount   float64 `json:"ovo_cash_amount"`
		OvoPointsAmount float64 `json:"ovo_points_amount"`
	}
)

type (
	PromoCodeUsageAffectedOrderMix struct {
		ID                             int64                         `json:"id"`
		PromoID                        int64                         `json:"promo_id"`
		PromoCodeUsageID               int64                         `json:"promo_code_usage_id"`
		OrderID                        int64                         `json:"order_id"`
		PaymentID                      int64                         `json:"payment_id"`
		CreateTime                     time.Time                     `json:"create_time"`
		UpdateTime                     time.Time                     `json:"update_time"`
		ServiceID                      int64                         `json:"service_id"`
		DiscountAmount                 float64                       `json:"discount_amount"`
		CashbackAmount                 float64                       `json:"cashback_amount"`
		GivenCashbackAmount            float64                       `json:"given_cashback_amount"`
		ShowCashbackAmount             bool                          `json:"show_cashback_amount"`
		ShowDiscountAmount             bool                          `json:"show_disocunt_amount"`
		PromoCodeUsageAffectedProducts map[int64]AffectedOrderDetail `json:"promo_code_usage_affected_products"`
		OrderDetail                    string                        `json:"order_detail"`
		Status                         int
	}
)
type AffectedOrderDetail struct {
	ProductID       int64                      `json:"product_id"`
	ProductPrice    float64                    `json:"product_price"`
	Promo           []AffectedOrderDetailPromo `json:"promo"`
	IsPowerBadge    bool                       `json:"is_power_badge"`
	IsOfficialStore bool                       `json:"is_official_store"`
}

type AffectedOrderDetailPromo struct {
	Code           string `json:"code"`
	Tier           string `json:"tier"`
	TieringItem    string
	DiscountAmount float64 `json:"discount_amount"`
	CashbackAmount float64 `json:"cashback_amount"`
	ShowTag        bool    `json:"show_tag"`
}

type PromoCode struct {
	ID                     int64     `json:"id"`
	PromoID                int64     `json:"promo_id"`
	Code                   string    `json:"code"`
	Status                 int       `json:"status"`
	PID                    int64     `json:"pid"`
	RefundAmount           float64   `json:"refund_amount"`
	RefundCount            int       `json:"refund_count"`
	Owner                  int64     `json:"owner"`
	CreateTime             time.Time `json:"create_time"`
	ExpiryTime             time.Time `json:"expiry_time"`
	IsQcCode               bool      `json:"is_qc_code"`
	CreateBy               int64     `json:"create_by"`
	UpdateBy               int64     `json:"update_by"`
	RefPaymentID           int64     `json:"ref_payment_id"`
	CountFinishedCashbacks int       `json:"count_finished_cashbacks"`
	WalletAmount           float64   `json:"wallet_amount"`
	TransactionID          string    `json:"transaction_id"`
	CompanyID              string    `json:"company_id"`
	OwnerEmail             string    `json:"owner_email"`
	IsCoupon               int       `json:"is_coupon"`
	IsRecycle              bool      `json:"is_recycle"`

	CouponTitle string `json:"coupon_title"`
}
