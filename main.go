package main

import (
	"fmt"
	"sort"
	"strings"
)

type Cart map[string]int

func main() {
    var varCart Cart

    cart := varCart.newCart()
    cart.addProduct("Pisang Hijau", 2)

    cart.addProduct("Semangka Kuning", 3)
	
    cart.addProduct("pisang merah", 3)

    cart.addProduct("Apel Merah", 1)
    cart.addProduct("Apel Merah", 4)
    cart.addProduct("Apel Merah", 2)

    // cart.deleteProduct("Semangka Kuning")
    
    cart.deleteProduct("Semangka Merah")
    
	cart.orderCart()
    // cart.showCart()

}

// order the products on Cart case-insensitively
func (cart Cart) orderCart() {
	keys := make([]string, 0)
	for k, _ := range cart {
		keys = append(keys, k)
	}
	// sort.Strings(keys) // Case sensitive sorting
	sort.Slice(keys, func(i, j int) bool { 
			return strings.ToLower(keys[i]) < strings.ToLower(keys[j])})

	for _, k := range keys {
		fmt.Println(k, cart[k])
	}
}

func (cart Cart) newCart() Cart {
    return make(Cart, 0)
}

func (cart Cart) addProduct(productCode string, quantity int){
    if _, exist := cart[productCode]; !exist {
        cart[productCode] = quantity
    } else {
        cart[productCode] = cart[productCode] + quantity
    }

    return
}

func (cart Cart) deleteProduct(productCode string) {
    if _, exist := cart[productCode]; exist {
        delete(cart, productCode)
    }

    return
}

func (cart Cart) showCart() {
    for key, value := range cart {
        fmt.Println(key," (", value,")")
    }

    return
}