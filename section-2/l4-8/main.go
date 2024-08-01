package main

// Определяем структуру MyStruct с полями On, Ammo и Power
type MyStruct struct {
	On    bool
	Ammo  int
	Power int
}

// Метод Shoot возвращает true, если можно сделать выстрел, иначе false
func (m *MyStruct) Shoot() bool {
	if !m.On {
		return false
	}
	if m.Ammo > 0 {
		m.Ammo--
		return true
	}
	return false
}

// Метод RideBike возвращает true, если можно покататься на велосипеде, иначе false
func (m *MyStruct) RideBike() bool {
	if !m.On {
		return false
	}
	if m.Power > 0 {
		m.Power--
		return true
	}
	return false
}

// Функция для тестирования методов структуры
func testStructFunc(m *MyStruct) {
	// Вызываем методы Shoot и RideBike несколько раз для проверки
	m.Shoot()
	m.RideBike()
	m.Shoot()
	m.RideBike()
}

func main() {
	// Создаем экземпляр структуры и присваиваем его указатель переменной testStruct
	testStruct := &MyStruct{
		On:    true,
		Ammo:  10,
		Power: 5,
	}

	// Экземпляр созданной вами структуры необходимо передать в качестве
	// аргумента функции testStruct, которая выполнит проверку соблюдения
	// всех условий задания.
	testStructFunc(testStruct)
}
