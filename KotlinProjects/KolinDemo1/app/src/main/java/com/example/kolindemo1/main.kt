package com.example.kolindemo1

fun main() {
//    val student = Student("jack", 19)
//    doStudy(student)
    val cellphone1 = Cellphone("Samsung", 1299.99)
    val cellphone2 = Cellphone("Samsung", 1299.99)
    println(cellphone1)
    println("cellphone1 eauals cellphone2 " + (cellphone1 == cellphone2))

    Singleton.singleTest()

    val list = listOf("apple", "banana", "orange", "pear", "grape")
    for (fruit in list) {
        println(fruit)
    }

    Thread(object : Runnable {
        override fun run() {
            println("Thread is running")
        }
    }).start()

    Thread(Runnable {
        println("Thread is running")
    }).start()

    Thread({
        println("Thread is running")
    }).start()

    Thread {
        println("Thread is running")
    }.start()

    doStudy(Student("jack", 20))

    doStudy(null)
}

fun largerNum(num1: Int, num2: Int): Int = if (num1 > num2) num1 else num2

fun doStudy(study: Study?) {
    if (study != null){
        study.readBooks()
    }
}

fun getScore(name: String) = when {
    name.startsWith("Tom") -> 86
    name == "jim" -> 77
    name == "jim" -> 77
    name == "jim" -> 77
    else -> 0
}

