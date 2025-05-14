package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Секундомер")

	// Поле для ввода времени
	input := widget.NewEntry()
	input.SetPlaceHolder("Введите время в секундах")

	// Прогрессбар
	progress := widget.NewProgressBar()
	progress.Min = 0.0
	progress.Max = 1.0

	// Текст, который будет показывать оставшееся время
	progressText := canvas.NewText("00:00", theme.ForegroundColor())
	progressText.Alignment = fyne.TextAlignCenter
	progressText.TextSize = 20

	// Кнопка для старта
	startBtn := widget.NewButton("Старт", func() {
		seconds, err := strconv.Atoi(input.Text)
		if err != nil || seconds <= 0 {
			progressText.Text = "Неверный ввод"
			progressText.Color = theme.Color(theme.ColorNameError)
			canvas.Refresh(progressText)
			return
		}

		// Устанавливаем прогресс в 100%
		progress.SetValue(1.0)
		progressText.Text = formatTime(seconds)
		progressText.Color = theme.ForegroundColor()
		canvas.Refresh(progress)
		canvas.Refresh(progressText)

		// Таймер для отсчёта
		go func() {
			for i := seconds; i >= 0; i-- {
				time.Sleep(time.Second)

				// Вычисляем процент оставшегося времени
				percent := float64(i) / float64(seconds)
				progress.SetValue(percent)

				// Обновление цвета и текста
				if percent > 0.66 {
					progressText.Color = theme.Color(theme.ColorNameSuccess) // Зелёный
				} else if percent > 0.33 {
					progressText.Color = theme.Color(theme.ColorNameWarning) // Оранжевый
				} else {
					progressText.Color = theme.Color(theme.ColorNameError) // Красный
				}

				// Обновляем текст с оставшимся временем
				progressText.Text = formatTime(i)
				canvas.Refresh(progressText)
				canvas.Refresh(progress)
			}

			// Когда таймер завершится
			progressText.Text = "Время вышло!"
			progressText.Color = theme.Color(theme.ColorNameError)
			canvas.Refresh(progressText)
		}()
	})

	// Размещение элементов в контейнере
	w.SetContent(container.NewVBox(
		input,
		startBtn,
		progress,
		progressText,
	))

	// Окно
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

// Форматирует время в виде 00:00
func formatTime(seconds int) string {
	minutes := seconds / 60
	seconds = seconds % 60
	return strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)
}
