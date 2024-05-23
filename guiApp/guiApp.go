package guiApp

import (
	"errors"
	"strconv"
	"strings"
	"wor/bst"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//------------------------------------------------------------------------------

var ErrConvertStrToInt error = errors.New("bst: String to int conversion errord")

//------------------------------------------------------------------------------

func convertStringToInt(str string) (int, error) {
	data, err := strconv.Atoi(str)
	if err != nil {
		return 0, ErrConvertStrToInt
	}
	return data, nil
}

// ------------------------------------------------------------------------------

func processDataInsert(valueInput string, tree *bst.BST) (string, int) {
	dataStrArr := strings.Split(valueInput, " ")
	errStr := ""
	var count int = 0

	numbers := make([]int, 0, len(dataStrArr))
	for _, dataStr := range dataStrArr {
		data, err := convertStringToInt(dataStr)
		if err != nil {
			errStr = errStr + err.Error() + " {" + dataStr + "}\n"
			continue
		}
		numbers = append(numbers, data)
	}

	for _, data := range numbers {
		errInsert := tree.Insert(data)
		if errInsert != nil {
			errStr = errStr + errInsert.Error() + " {" + strconv.Itoa(data) + "}\n"
			continue
		}
		count++
	}
	return errStr, count
}

//------------------------------------------------------------------------------

func processDataRemove(removeInput string, tree *bst.BST) (string, int) {
	dataStrArr := strings.Split(removeInput, " ")
	errStr := ""
	var count int = 0

	numbers := make([]int, 0, len(dataStrArr))
	for _, dataStr := range dataStrArr {
		data, err := convertStringToInt(dataStr)
		if err != nil {
			errStr = errStr + err.Error() + " {" + dataStr + "}\n"
			continue
		}
		numbers = append(numbers, data)
	}

	for _, data := range numbers {
		errInsert := tree.Remove(data)
		if errInsert != nil {
			errStr = errStr + errInsert.Error() + " {" + strconv.Itoa(data) + "}\n"
			continue
		}
		count++
	}
	return errStr, count
}

//------------------------------------------------------------------------------

func processDataReplace(replaceOldInput string, replaceNewInput string, tree *bst.BST) (string, int) {
	dataStrOld := replaceOldInput
	dataStrNew := replaceNewInput

	errStr := ""
	count := 0

	for {
		dataOld, errOld := convertStringToInt(dataStrOld)
		dataNew, errNew := convertStringToInt(dataStrNew)

		if errOld != nil || errNew != nil {
			if errOld != nil {
				errStr = errStr + errOld.Error() + " {" + strconv.Itoa(dataOld) + "}\n"
			}

			if errNew != nil {
				errStr = errStr + errNew.Error() + " {" + strconv.Itoa(dataNew) + "}\n"
			}
			break
		}

		err := tree.Replace(dataOld, dataNew)
		if err != nil {
			errStr = errStr + err.Error() + "\n"
			break
		} else {
			count++
			break
		}
	}
	return errStr, count
}

//------------------------------------------------------------------------------

func StartGUI(tree *bst.BST) {
	a := app.New()
	a.Settings().SetTheme(dejaVuSansMonoTheme{})
	w := a.NewWindow("BST GUI")

	w.Resize(fyne.NewSize(500, 900))

	treeOutput := widget.NewLabel(tree.TreeString())

	sizeOutput := widget.NewLabel(strconv.Itoa(tree.Size()))
	preorderOutput := widget.NewLabel(tree.PreorderStr())
	inorderOutput := widget.NewLabel(tree.InorderStr())
	postorderOutput := widget.NewLabel(tree.PostorderStr())
	statusOutput := widget.NewLabel(tree.PostorderStr())
	errorOutput := widget.NewLabel("")

	valueInput := widget.NewEntry()
	valueInput.SetPlaceHolder("Enter value to insert")

	removeInput := widget.NewEntry()
	removeInput.SetPlaceHolder("Enter value to remove")

	replaceOldInput := widget.NewEntry()
	replaceOldInput.SetPlaceHolder("Enter old value to replace")

	replaceNewInput := widget.NewEntry()
	replaceNewInput.SetPlaceHolder("Enter new value")

	//------------------------------------------------------------------------------

	insertButton := widget.NewButton("Insert",
		func() {
			errStr, count := processDataInsert(valueInput.Text, tree)

			errorOutput.SetText(errStr)
			treeOutput.SetText(tree.TreeString())
			sizeOutput.SetText(strconv.Itoa(tree.Size()))
			preorderOutput.SetText(tree.PreorderStr())
			inorderOutput.SetText(tree.InorderStr())
			postorderOutput.SetText(tree.PostorderStr())
			statusOutput.SetText("Number of successful inserted nodes: " + strconv.Itoa(count) + " out of " + strconv.Itoa(len(strings.Split(valueInput.Text, " "))))
		})

	//------------------------------------------------------------------------------

	removeButton := widget.NewButton("Remove",
		func() {
			errStr, count := processDataRemove(removeInput.Text, tree)

			errorOutput.SetText(errStr)
			treeOutput.SetText(tree.TreeString())
			sizeOutput.SetText(strconv.Itoa(tree.Size()))
			preorderOutput.SetText(tree.PreorderStr())
			inorderOutput.SetText(tree.InorderStr())
			postorderOutput.SetText(tree.PostorderStr())
			statusOutput.SetText("Number of successful removed nodes: " + strconv.Itoa(count) + " out of " + strconv.Itoa(len(strings.Split(removeInput.Text, " "))))
		})

	//------------------------------------------------------------------------------

	replaceButton := widget.NewButton("Replace",
		func() {
			errStr, count := processDataReplace(replaceOldInput.Text, replaceNewInput.Text, tree)

			errorOutput.SetText(errStr)
			treeOutput.SetText(tree.TreeString())
			sizeOutput.SetText(strconv.Itoa(tree.Size()))
			preorderOutput.SetText(tree.PreorderStr())
			inorderOutput.SetText(tree.InorderStr())
			postorderOutput.SetText(tree.PostorderStr())
			statusOutput.SetText("Number of successful replaced nodes: " + strconv.Itoa(count) + " out of 1")
		})

	//------------------------------------------------------------------------------

	clearButton := widget.NewButton("Clear",
		func() {
			tree.Root(nil)
			tree.Size(0)
			errorOutput.SetText("")
			treeOutput.SetText(tree.TreeString())
			sizeOutput.SetText(strconv.Itoa(tree.Size()))
			preorderOutput.SetText(tree.PreorderStr())
			inorderOutput.SetText(tree.InorderStr())
			postorderOutput.SetText(tree.PostorderStr())
			statusOutput.SetText("")

		})

	//------------------------------------------------------------------------------

	treeContainer := container.NewVScroll(treeOutput)
	treeContainer.SetMinSize(fyne.NewSize(400, 400))

	//------------------------------------------------------------------------------

	displayForm := widget.NewForm(

		widget.NewFormItem("Size:", sizeOutput),
		widget.NewFormItem("Preorder:", preorderOutput),
		widget.NewFormItem("Inorder:", inorderOutput),
		widget.NewFormItem("Postorder:", postorderOutput),
		widget.NewFormItem("Status:", statusOutput),
		widget.NewFormItem("ERROR:", errorOutput),
	)

	displayContainer := container.NewVSplit(treeContainer, container.NewHScroll(container.NewVScroll(displayForm)))

	//------------------------------------------------------------------------------

	insertContainer := container.NewVBox(valueInput, insertButton)

	removeContainer := container.NewVBox(removeInput, removeButton)

	replaceContainer := container.NewVBox(
		container.NewHSplit(replaceOldInput, replaceNewInput),
		replaceButton,
	)

	interactionForm := widget.NewForm(
		widget.NewFormItem("Insert Value:", insertContainer),
		widget.NewFormItem("Remove Value:", removeContainer),
		widget.NewFormItem("Replace Value:", replaceContainer),
		widget.NewFormItem("Clear BST:", clearButton),
	)

	//------------------------------------------------------------------------------

	w.SetContent(container.NewVSplit(
		displayContainer,
		interactionForm,
	))

	w.ShowAndRun()
}
