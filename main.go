package main

import (
	"strconv"
	"syscall/js"
)

func click(id string) {
	//idから対称のオブジェクトを取得
	object := js.Global().Get("document").Call("getElementById", id)
	//操作内容
	action := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//画像の差し替え
		if object.Get("alt").String() == "image2" {
			object.Set("alt", string("image1"))
			object.Set("src", string("image/image1.png"))
		} else {
			object.Set("alt", string("image2"))
			object.Set("src", string("image/image2.png"))
		}
		return nil
	})
	//アクションの呼び出し
	object.Call("addEventListener", "click", action)
}

func mouseIn(id string) {
	object := js.Global().Get("document").Call("getElementById", id)
	action := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//画像の差し替え
		if object.Get("alt").String() == "image2" {
			object.Set("alt", string("image1"))
			object.Set("src", string("image/image1.png"))
		} else {
			object.Set("alt", string("image2"))
			object.Set("src", string("image/image2.png"))
		}
		return nil
	})
	//画像にマウスカーソルが入った瞬間動作する
	object.Call("addEventListener", "mouseover", action)
}

func mouseOver(id string) {
	object := js.Global().Get("document").Call("getElementById", id)
	action := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//画像の差し替え
		if object.Get("alt").String() == "image2" {
			object.Set("alt", string("image1"))
			object.Set("src", string("image/image1.png"))
		} else {
			object.Set("alt", string("image2"))
			object.Set("src", string("image/image2.png"))
		}
		return nil
	})
	//画像からマウスカーソルが出た瞬間動作する
	object.Call("addEventListener", "mouseout", action)
}

// func touch(id string) {
// 	object := js.Global().Get("document").Call("getElementById", id)
// 	action := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		if object.Get("alt").String() == "image2" {
// 			object.Set("alt", string("image1"))
// 			object.Set("src", string("image/image1.png"))
// 		} else {
// 			object.Set("alt", string("image2"))
// 			object.Set("src", string("image/image2.png"))
// 		}
// 		return nil
// 	})
// 	//スマホでタッチしたとき
// 	object.Call("addEventListener", "touchstart", action)
// }

func text_change(id string, id2 string) {
	obj := js.Global().Get("document").Call("getElementById", id)
	text := js.Global().Get("document").Call("getElementById", id2)
	tmp := text.Get("textContent").String()
	num, _ := strconv.Atoi(tmp)
	//
	action := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		num++
		tmp = strconv.Itoa(num)
		text.Set("innerHTML", tmp)
		return nil
	})
	obj.Call("addEventListener", "click", action)
}

func get_mouthXY(id string, output string) {
	obj := js.Global().Get("document").Call("getElementById", id)
	out := js.Global().Get("document").Call("getElementById", output)

	//-----------------------------------------------------------------------------------
	move := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		mx := args[0].Get("clientX").Float()
		my := args[0].Get("clientY").Float()

		out.Set("innerHTML", "X:"+strconv.FormatFloat(mx, 'f', -1, 32)+" Y:"+strconv.FormatFloat(my, 'f', -1, 32))

		return nil
	})
	//-----------------------------------------------------------------------------------
	obj.Call("addEventListener", "mousemove", move)

}

func main() {

	//イベントの管理
	click("image1")
	text_change("image1", "text1")
	mouseIn("image2")
	mouseOver("image3")
	get_mouthXY("image4", "text2")
	//touch("image3")
	// プログラムが終了しないよう、チャネルの待ち受けをする
	<-make(chan struct{}, 0)

}
