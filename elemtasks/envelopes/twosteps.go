package main

import "fmt"

/*Проверка конвертов в 2 шага + если нужно, меняю их местами и повторяю 2 шага.*/
func check(env1, env2 *envelope) string {
	if ok := step1(env1, env2); ok {
		return fmt.Sprintf("[%s] может поместиться в [%s].", env2.name, env1.name)
	}
	fmt.Printf("Шаг 1 неудача, поворачиваю [%s] на 90 градусов и снова пытаюсь вложить в [%s].\n", env2.name, env1.name)

	if ok := step2(env1, env2); ok {
		return fmt.Sprintf("Поворот помог и теперь [%s] таки можно вложить в [%s].", env2.name, env1.name)
	}
	fmt.Printf("Поворот не помог, меняю конверты местами и теперь пытаюсь вложить [%s] в [%s].\n", env1.name, env2.name)

	if ok := step1(env2, env1); ok {
		return fmt.Sprintf("[%s] может поместиться в [%s].", env1.name, env2.name)
	}
	fmt.Printf("2-ой круг, шаг 1 > неудача, поворачиваю [%s] на 90 градусов и снова пытаюсь вложить в [%s].\n", env1.name, env2.name)

	if ok := step2(env2, env1); ok {
		return fmt.Sprintf("Теперь [%s] может поместиться в [%s].", env1.name, env2.name)
	}

	/*Если дошел до этой точки.*/
	return fmt.Sprintf("Неудача!) Ни [%s] не может быть вложен в [%s], ни [%s] не может быть "+
		"вложен в [%s].", env1.name, env2.name, env2.name, env1.name)
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
