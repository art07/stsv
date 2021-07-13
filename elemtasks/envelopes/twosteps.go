package main

import "fmt"

/*Проверка конвертов в 2 шага + если нужно, меняю их местами и повторяю 2 шага.*/
func check() (answer string) {
	ok := step1(&env1, &env2)
	if ok {
		return "Конверт 2 может поместиться в конверте 1."
	}

	fmt.Println("Шаг 1 неудача, поворачиваю конверт 2 на 90 градусов и снова пытаюсь вложить в конверт 1.")

	ok = step2(&env1, &env2)
	if ok {
		return "Поворот помог и теперь конверт 2 может поместиться в конверте 1."
	}

	fmt.Println("Не помогло, меняю конверты местами и теперь пытаюсь вложить конверт 1 в конверт 2.")

	ok = step1(&env2, &env1)
	if ok {
		return "Конверт 1 может поместиться в конверте 2."
	}

	fmt.Println("Повторный шаг 1 неудача, поворачиваю конверт 1 на 90 градусов и снова пытаюсь вложить в конверт 2.")

	ok = step2(&env2, &env1)
	if ok {
		return "Теперь конверт 1 может поместиться в конверте 2."
	}

	/*Если дошел до этой точки.*/
	return "Неудача!) Ни конверт 1 не может быть вложен в конверт 2, ни конверт 2 не может быть вложен в конверт 1."
}

func step1(env1, env2 *envelope) (ok bool) {
	if env2.side1 < env1.side2 && env2.side2 < env1.side1 {
		ok = true
	}
	return
}

func step2(env1, env2 *envelope) (ok bool) {
	if env2.side2 < env1.side2 && env2.side1 < env1.side1 {
		ok = true
	}
	return
}
