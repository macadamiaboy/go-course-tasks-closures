// Задание 3: Почему WaitGroup передаётся параметром, а не живёт снаружи?
//
// Ты уже видел в своём коде такое:
//
//   var wg sync.WaitGroup  // глобально, снаружи функции
//
//   func worker(jobs <-chan int, results chan<- int) {
//       defer wg.Done()  // захватывает глобальный wg
//   }
//
// Почему это плохо? Разберём по шагам.
//
// WaitGroup - это счётчик "сколько горутин ещё работают".
// Когда wg живёт глобально, у этого три проблемы:
//
//   1. Нельзя запустить пул дважды в одной программе.
//      После первого wg.Wait() счётчик равен 0.
//      Если стартовать снова - goroutine panic при wg.Add(-1) (Done без Add).
//
//   2. Код трудно читать.
//      Непонятно кто и где трогает wg. Функция worker зависит от
//      глобального состояния, которое не видно в её сигнатуре.
//
//   3. Тесты писать невозможно.
//      Нельзя протестировать worker изолированно - у него скрытая зависимость.
//
// ПРАВИЛЬНО: передавать wg явно как параметр:
//
//   func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
//   //                                              ^^^^^^^^^^^^^^^^^^
//   //           явно видно: этот worker обязан вызвать wg.Done()
//       defer wg.Done()
//   }
//
// Зависимость становится видимой. wg создаётся там, где нужен - в main() или функции.
//
// ──────────────────────────────────────────────────────────────────────────────
//
// Задача: реализуй worker pool ПРАВИЛЬНО - с wg как параметром.
//
// Напиши функцию worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup),
// которая:
//   - при завершении вызывает defer wg.Done()
//   - читает задачи из jobs через range
//   - для каждой задачи j формирует строку fmt.Sprintf("воркер %d: %d^2 = %d", id, j, j*j)
//   - пишет её в results
//
// В main():
//   1. Создай WaitGroup ЛОКАЛЬНО: var wg sync.WaitGroup
//   2. Создай каналы jobs и results (буферизованные, размер 10)
//   3. Запусти 3 воркера - передай &wg каждому
//   4. Закинь задачи 1..10 в jobs и закрой jobs
//   5. Запусти горутину: wg.Wait() → close(results)
//   6. Выведи все результаты из results
//
// Ожидаемый вывод (порядок строк разный каждый раз):
//   воркер 2: 5^2 = 25
//   воркер 1: 1^2 = 1
//   воркер 3: 3^2 = 9
//   ... и т.д.
//
// Проверь что код работает если вызвать запуск пула дважды подряд.
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"sync"
)

// TODO: напиши функцию worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup)
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		results <- fmt.Sprintf("воркер %d: %d^2 = %d", id, job, job*job)
	}
}

func main() {
	// TODO: объяви WaitGroup ЛОКАЛЬНО здесь
	var wg sync.WaitGroup
	// 	var wg2 sync.WaitGroup

	// TODO: создай каналы jobs и results
	jobs := make(chan (int), 10)
	results := make(chan (string), 10)

	// TODO: запусти 3 воркера, передавая &wg
	for i := range 3 {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// TODO: закинь задачи 1..10 в jobs и закрой jobs
	for i := range 10 {
		// wg2.Add(1)
		// go func() {
		//	defer wg2.Done()
		jobs <- i
		//}()
	}
	close(jobs)
	/*
		go func() {
			wg2.Wait()
			close(jobs)
		}()
	*/
	// TODO: запусти горутину которая после wg.Wait() закрывает results
	go func() {
		wg.Wait()
		close(results)
	}()

	// TODO: выведи все результаты из results через range
	for i := range results {
		fmt.Println(i)
	}
}
