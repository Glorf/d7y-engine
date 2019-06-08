import "fmt"
import "strings"
import "os"


type OrderContent struct {
    From     string `json:from`
    To       string `json:to`
    UnitType string `json:unitType`
}

type Order struct {
    OrderType    string `json:orderType`
    Player       string `json:player`
    Location     string `json:location`
    UnitType     string `json:unitType`
    OrderContent `json:orderContent`
}




func main() {
    //email := "F Par S A Mar-Bur"
    email := os.Args[1]
    // var resp = http.Get("api url")
    // var provinceList []string
    // json.Unmarshall([]byte(resp), &provinceList)
    // var provinceList = [3]string{"Par", "Mar", "Bur"}

    orders := parseOrders(email, "test")

    fmt.Printf("%v", orders)


}


// example orders
// move: A Par-Bur
// support: F Par S A Mar-Bur


func parseOrders(message string, username string) []Order {
    ordersStringList := strings.Split(strings.ToLower(strings.ReplaceAll(message, "-", " ")), "\n")
    var ordersList []Order
    for index := range ordersStringList {
        splitOrder := strings.Fields(ordersStringList[index])
        if len(splitOrder) == 3 {
            if len(splitOrder[0]) != 1 || len(splitOrder[1]) != 3 || len(splitOrder[2]) != 3 {
                return nil
            }
            if !strings.Contains("af", splitOrder[0]) {
                return nil
            }
            order := Order{
                Player:    username,
                OrderType: "M",
                Location:  splitOrder[1],
                UnitType:  splitOrder[0],
                OrderContent: OrderContent{
                    From: splitOrder[1],
                    To:   splitOrder[2],
                },
            }
            ordersList = append(ordersList, order)
        } else if len(splitOrder) == 6 {
            if len(splitOrder[0]) != 1 || len(splitOrder[1]) != 3 || len(splitOrder[2]) != 1 ||
               len(splitOrder[3]) != 1 || len(splitOrder[4]) != 3 || len(splitOrder[5]) != 3 {
                return nil
            }
            if !(strings.Contains("af", splitOrder[0]) && strings.Contains("sc", splitOrder[2]) &&
                 strings.Contains("af", splitOrder[3])) {
                return nil
            }
            order := Order{
                Player:    username,
                OrderType: splitOrder[2],
                Location:  splitOrder[1],
                UnitType:  splitOrder[0],
                OrderContent: OrderContent{
                    UnitType: splitOrder[3],
                    From:     splitOrder[4],
                    To:       splitOrder[5],
                },
            }
            ordersList = append(ordersList, order)
        } else {
            return nil
        }
    }
    return ordersList
}

