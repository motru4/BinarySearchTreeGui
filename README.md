# Project "BinarySearchTreeGui"

## Description
"BinarySearchTreeGui" is a desktop application that implements visualization and editing of a binary search tree in Golang using the Fyne library.

## Project Contents
The project contains the following files and directories:

- `wor.exe`: executable file of the application.
- `tree.png`: application icon.
- `main.go`: main program file.
- `go.sum` and `go.mod`: Go module files containing dependency information.
- `guiApp/`: directory containing files for the GUI application.
  - `bundled.go`: file containing bundled resources for GUI.
  - `guiApp.go`: main GUI application file.
  - `theme.go`: file defining the GUI application theme.
- `consoleApp/consoleApp.go`: console application file.
- `bst/`: directory containing files for the binary search tree.
  - `bst_interface.go`: file defining the binary search tree interface.
  - `bst_test.go`: file containing tests for the binary search tree.
  - `node.go`: file defining the tree node structure.
  - `tree.go`: file defining the tree structure.
- `assets/DejaVuSansMono.ttf`: font file used in the application.

## Installation and Launch
To install and launch the project, you will need Go and Fyne:

- Go: Installation instructions https://go.dev/doc/install
- Fyne: Installation instructions https://docs.fyne.io/started

  After installing these dependencies, you can clone the repository and run `main.go`.

To run the executable file, nothing is required:

- `wor.exe`: for Windows.
- `wor_linux`: for Linux. Will be added later!
- `wor_macos`: for MacOS. Will be added later!

## Usage Example

After installing and launching the application, you will see a graphical user interface where you can edit a binary search tree. Here is a usage example:

![img_wor](https://github.com/motru4/BinarySearchTreeGui/assets/92480080/fbaea62b-667d-47d8-ada0-9a15e1e53150)

1. Launch the application by running `wor.exe`.
2. Add a node with value X by entering “X” in the input field and pressing the “Insert” button. Add a few more nodes by repeating this process or by entering insertion values separated by a space.
3. Remove a node with value X by entering “X” in the input field and pressing the “Remove” button. To remove multiple nodes, you can specify node values separated by a space, as when adding.
4. Change a node from an old value to a new value by entering the old and new values in the corresponding fields and pressing the “Replace” button.
5. Clear the tree by pressing the “Clear” button.

Description of the output data:
1. Size: Tree size.
2. Preorder: Tree representation using pre-order traversal.
3. Inorder: Tree representation using in-order traversal.
4. Postorder: Tree representation using post-order traversal.
5. Status: Number of successful operations .
6. ERROR: Various kinds of errors.

## License
This project is distributed under the MIT license. Details can be found in the LICENSE file.

---

# Проект "BinarySearchTreeGui"

## Описание
"BinarySearchTreeGui" - это настольное приложение, реализующее визуализацию и редактирование бинарное дерево поиска на языке Golang с использованием библиотеки Fyne.

## Содержание проекта
Проект содержит следующие файлы и каталоги:

- `wor.exe`: исполняемый файл приложения.
- `tree.png`: иконка приложения.
- `main.go`: основной файл программы.
- `go.sum` и `go.mod`: файлы модулей Go, содержащие информацию о зависимостях проекта.
- `guiApp/`: каталог, содержащий файлы для GUI приложения.
  - `bundled.go`: файл, содержащий встроенные ресурсы для GUI.
  - `guiApp.go`: основной файл GUI приложения.
  - `theme.go`: файл, определяющий тему GUI приложения.
- `consoleApp/consoleApp.go`: файл консольного приложения. 
- `bst/`: каталог, содержащий файлы для бинарного дерева поиска.
  - `bst_interface.go`: файл, определяющий интерфейс бинарного дерева поиска.
  - `bst_test.go`: файл, содержащий тесты для бинарного дерева поиска.
  - `node.go`: файл, определяющий структуру узла дерева.
  - `tree.go`: файл, определяющий структуру дерева.
- `assets/DejaVuSansMono.ttf`: файл шрифта, используемый в приложении.

## Установка и запуск
Для установки и запуска проекта вам потребуется Go и Fyne:

- Go: Инструкция по установке https://go.dev/doc/install
- Fyne: Инструкция по установке https://docs.fyne.io/started

  После установки этих зависимостей вы можете клонировать репозиторий и запустить `main.go`.

Для запуска исполняемого файла ничего не требуется:

- `wor.exe`: для Windows.
- `wor_linux`: для Linux. Будет добавлен позднее!
- `wor_macos`: для MacOS. Будет добавлен позднее!

## Пример использования

После установки и запуска приложения вы увидите графический интерфейс пользователя, где можно вводить редактировать бинарное дерево поиска. Вот пример использования:

![img_wor](https://github.com/motru4/BinarySearchTreeGui/assets/92480080/fbaea62b-667d-47d8-ada0-9a15e1e53150)

1. Запустите приложение, выполнив `wor.exe`.
2. Добавьте узел со значением X, введя “X” в поле для ввода и нажав кнопку “Insert”. Добавьте еще несколько узлов, повторив этот процесс или введя значения для вставки через пробел.
3. Удалите узел со значением X, введя “X” в поле для ввода и нажав кнопку “Remove”. Для удаления нескольких узлов, можно как при добавлении указывать значение узлов через пробел.
4. Измените узел со старого значением на новое значение, введя старое и новое значения в соответствующие поля и нажав кнопку “Replace”.
5. Очистите дерево нажав кнопку “Clear”.

Описание выводимых данных:
1. Size: Размер дерева.
2. Preorder: Представление дерева использующее прямой обход.
3. Inorder: Представление дерева использующее центрированный обход.
4. Postorder: Представление дерева использующее обратный обход.
5. Status: Количество удачно проведенных операций .
6. ERROR: Разного рода ошибки.

## Лицензия
Этот проект распространяется под лицензией MIT. Подробности можно найти в файле LICENSE.
