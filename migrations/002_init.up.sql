CREATE TABLE sessions (
  session_id char(128) PRIMARY KEY,
  user_id BIGINT,
  questions JSONB,
  user_answers INT[],
  expired_at TIMESTAMP
);

CREATE TABLE questions (
  question_id SERIAL PRIMARY KEY,
  question_text TEXT NOT NULL,
  question_options JSONB NOT NULL,
  answer INT NOT NULL,
  explanation TEXT
);

INSERT INTO questions (question_text, question_options, answer, explanation)
VALUES
('Что такое горутина (goroutine) в Go?', '["Класс объектов", "Поток выполнения", "Интерфейс", "Группа функций"]', 1, 'Это легковесные потоки, которые реализуют конкурентное программирование. Они управляются рантаймом, а не операционной системой'),
('Каким образом проверить ошибку в Go?', '["if err != nil { }", "err := checkError()", "if error { }", "if err == nil { }"]', 0, 'Нужно сначала вернуть из функции переменную err, а потом проверить ее на nil'),
('Что такое defer в Go?', '["Запись данных", "Выполнение после возврата", "Оператор итерации", "Отложенное выполнение"]', 3, 'Ключевое слово defer используется для отложенного выполнения функции/метода до тех пор, пока текущая функция не завершится'),
('Каким образом происходит передача параметров в функции в Go?', '["По значению (pass by value)", "По ссылке (pass by reference)", "По указателю", "По имени параметра"]', 0, 'Всякий раз, когда мы передаем аргумент в функцию, функция получает копию первоначального значения'),
('Что такое слайс (slice) в Go?', '["Указатель на массив", "Динамический массив", "Класс для работы с файлами", "Массив фиксированного размера"]', 1, 'Это массив неопределённой длины'),
('Что такое интерфейс в Go?', '["Набор методов", "Имя типа данных", "Указатель на функцию", "Описание структуры"]', 0, 'Набор методов'),
('Какие управляющие конструкции есть в Go?', '["if, else, for, switch", "class, extends, implements", "begin, end, else, loop", "while, do-while, switch"]', 0, 'if, else, for, switch'),
('Каким образом происходит обработка ошибок в Go?', '["Исключения", "Возврат ошибки в функции", "Ловушка для ошибок", "Паника"]', 1, 'Ошибки - значения, возвращаемые функциями'),
('Что такое метод в Go?', '["Функция, связанная с определенным типом", "Глобальная функция", "Макрос", "Синтаксическая конструкция"]', 1, 'В Go нет классов, но существуют структуры с методами. Метод — это функция с дополнительным аргументом, который указывается в скобках между func и названием функции'),
('Go - императивный или декларативный?', '["не знаю", "декларативный", "императивный и декларативный", "императивный"]', 3, 'Императивное программирование — это описание того, как ты делаешь что-то (т.е. конкретно описываем необходимые действия для достижения определенного результата), а декларативное — то, что ты делаешь'),
('Что представляют из себя строки в go?', '["не знаю", "массив байт", "тип данных, предназначенный для работы с текстом", "любые текстовые данные"]', 1, 'это последовательность байт'),
('Что получится, если разделить int на 0, float на 0?', '["паника в обоих случаях", "невозможно в обоих случаях", "ошибка/бесконечность", "ноль/ноль"]', 2, 'Деление int на 0 в go невозможно и вызовет ошибку компилятора, а деление float на 0 дает в своем результате бесконечность.'),
('Как создать новую горутину в Go?', '["go func()", "new goroutine", "goroutine func()", "start go func"]', 0, 'Для создания новой горутины в Go используется ключевое слово go перед вызовом функции.'),
('Как вызвать метод defer в Go?', '["delay func()", "defer function()", "defer call()", "defer func()"]', 3, 'Для отложенного выполнения функции в Go используется ключевое слово defer перед вызовом функции.'),
('Как объявить константу в Go?', '["const x = 10", "constant x = 10", "let x = 10", "var x = 10"]', 0,
'Константы в Go объявляются с использованием ключевого слова const.'),
('Как проверить длину слайса в Go?', '["size(slice)", "length(slice)", "len(slice)", "count(slice)"]', 2, 'Для проверки длины слайса в Go используется функция len().'),
('Как объявить структуру в Go?', '["struct Struct {}", "type Struct struct {}", "struct {} Struct", "define Struct {}"]', 1, 'В Go структуры объявляются с использованием ключевых слов type и struct.'),
('Как сделать цикл for в Go?', '["while i < 10 {}", "let (i < 10) {}", "for i in 0..10 {}", "for i := 0; i < 10; i++ {}"]', 3, 'В Go цикл for используется для итераций и объявляется как for i := 0; i < 10; i++ {}.'),
('Как объявить map в Go?', '["make(map[string]int)", "create(map[string]int)", "new(map[string]int)", "map[string]int{}"]', 0, 'В Go карты объявляются с использованием функции make().'),
('Как создать канал в Go?', '["create(chan int)", "make(chan int)", "new(chan int)", "chan int{}"]', 1, 'Для создания канала в Go используется функция make().'),
('Что делает оператор select в Go?', '["Выбирает случайный канал", "Параллельно выполняет функции", "Синхронизирует выполнение горутин", "Выполняет операции с несколькими каналами"]', 3, 'select позволяет go-процедуре находиться в ожидании нескольких операций передачи данных. select блокируется до тех пор, пока один из его блоков case не будет готов к запуску, а затем выполняет этот блок. Если сразу несколько блоков могут быть запущены, то выбирается рандомный'),
('Как инициализировать слайс с длиной 5 в Go?', '["make([]int, 5)", "create([]int, 5)", "new([]int, 5)", "[]int{5}"]', 0, 'Для инициализации слайса с длиной 5 в Go используется функция make([]int, 5).'),
('Что делает функция panic() в Go?', '["Восстанавливает выполнение программы", "Печатает сообщение об ошибке", "Завершает программу", "Останавливает выполнение горутины"]', 2, 'Функция panic() в Go вызывает паническое завершение программы.'),
('Что возвращает функция recover() в Go?', '["Ошибку", "Паническое значение", "Значение nil", "Статус выполнения"]', 1, 'Функция recover() в Go возвращает паническое значение, если она вызывается в отложенной функции после panic().'),
('Как в Go избежать состояния гонки?', '["Использовать мьютексы", "Использовать каналы", "Использовать атомарные операции", "Все вышеперечисленное"]', 3, 'В Go можно избежать состояния гонки, используя мьютексы, каналы и атомарные операции.'),
('Что делает функция runtime.Gosched() в Go?', '["Завершает текущую горутину", "Усыпляет текущую горутину", "Передает управление другой горутине", "Останавливает все горутины"]', 3, 'Функция runtime.Gosched() передает управление другой горутине, позволяя планировщику выполнить другие задачи.'),
('Как в Go реализуется наследование?', '["Использование интерфейсов", "Встраивание структур", "Наследование классов", "Все вышеперечисленное"]', 1, 'В Go наследование реализуется через встраивание структур.'),
('Какую роль играет интерфейс в Go?', '["Определяет методы, которые должен реализовать тип", "Объявляет поля структуры", "Определяет пакет", "Объявляет переменную"]', 0, 'Интерфейс в Go определяет методы, которые должен реализовать тип.'),
('Что такое goroutine leak?', '["Потеря горутины", "Завершение горутины", "Утечка горутины", "Ошибка в горутине"]', 2, 'Goroutine leak — это утечка горутины, которая продолжает работать, когда она уже не нужна.'),
('Какой из подходов лучше всего использовать для параллельной обработки данных в Go?', '["Структуры и интерфейсы", "Мьютексы и условные переменные", "Стандартные функции и методы", "Горутины и каналы"]', 3, 'В Go для параллельной обработки данных лучше всего использовать горутины и каналы.'),
('Что такое context в Go?', '["Объект для хранения глобальных переменных", "Объект для передачи метаданных между горутинами", "Объект для управления временем выполнения задач", "Объект для организации структуры данных"]', 2, 'Context в Go используется для управления временем выполнения задач и передачи метаданных между горутинами.'),
('Как в Go реализовать чистую архитектуру?', '["Разделение логики на слои", "Использование интерфейсов для взаимодействия между слоями", "Использование зависимостей, инъекций и тестов", "Все вышеперечисленное"]', 3, 'В Go чистая архитектура реализуется через разделение логики на слои, использование интерфейсов и зависимостей, инъекций и тестов.'),
('Как определить deadlock в Go программе?', '["Программа завершает работу с ошибкой", "Все горутины заблокированы", "Программа зависает", "Все вышеперечисленное"]', 3, 'Deadlock в Go программе проявляется, когда все горутины заблокированы и программа зависает или завершает работу с ошибкой.'),
('Что такое race condition и как его избежать в Go?', '["Конкуренция за ресурсы между горутинами", "Проблемы с памятью", "Неправильное использование функций", "Ошибки при работе с интерфейсами"]', 0, 'Race condition — это конкуренция за ресурсы между горутинами. Его можно избежать, используя мьютексы, каналы и атомарные операции.'),
('Как использовать пакет sync в Go для синхронизации горутин?', '["sync.Mutex", "sync.WaitGroup", "sync.Cond", "Все вышеперечисленное"]', 3, 'Пакет sync в Go предоставляет такие примитивы синхронизации, как sync.Mutex, sync.WaitGroup и sync.Cond.'),
('Как реализовать шаблон проектирования "Фабрика" в Go?', '["Использовать функции-конструкторы", "Использовать методы структур", "Использовать интерфейсы", "Использовать пакеты"]', 0, 'В Go шаблон проектирования "Фабрика" можно реализовать с помощью функций-конструкторов.'),
('Как в Go работает garbage collection?', '["Оптимизирует работу программы", "Управляет выделением и освобождением памяти вручную", "Освобождает память, занимаемую ненужными объектами", "Автоматически исправляет ошибки"]', 2, 'Garbage collection в Go автоматически освобождает память, занимаемую ненужными объектами.'),
('Что такое pattern matching (сопоставление значения с шаблоном) в Go?', '["Сопоставление с шаблоном", "Ответы a, c", "Использование регулярных выражений", "Go не поддерживает pattern matching"]', 3, 'Go не поддерживает pattern matching, который присутствует в других языках, таких как Haskell или Erlang.'),
('Как в Go реализовать обобщенные типы?', '["Использовать интерфейсы", "Использовать пустые интерфейсы", "Использовать параметры типов", "Go не поддерживает обобщенные типы"]', 3, 'На данный момент Go не поддерживает обобщенные типы напрямую.'),
('Каким образом проверить ошибку в Go?', '["if err != nil { }", "err := checkError()", "if error { }", "if err == nil { }"]', 0, 'Нужно сначала вернуть из функции переменную err, а потом проверить ее на nil'),


('Какой из методов соответствует интерфейсу io.Reader в Go?', '["Write(p []byte) (n int, err error)", "Read(p []byte) (n int, err error)", "Close() error", "Open(name string) (file *File, err error)"]', 1, 'Метод Read(p []byte) (n int, err error) соответствует интерфейсу io.Reader, который позволяет считывать данные из источника в буфер p.'),
('Что такое метод Receiver в Go?', '["Метод, который возвращает значение", "Метод, который изменяет структуру", "Метод, привязанный к типу", "Метод, который вызывает другой метод"]', 2, 'Метод Receiver в Go — это метод, который привязан к конкретному типу, и этот тип может быть либо struct, либо любой другой пользовательский тип.'),
('Какое ключевое слово используется для определения нового типа в Go?', '["type", "struct", "interface", "package"]', 0, 'Ключевое слово type используется для определения нового типа в Go.'),
('Какое значение принимает неинициализированная переменная int в Go?', '["nil", "1", "0", "-1"]', 2, 'Неинициализированная переменная int принимает значение 0 в Go.'),
('Как вызвать метод, привязанный к структуре, в Go?', '["instance.method()", "method(instance)", "instance->method()", "instance:method()"]', 0, 'Метод, привязанный к структуре, вызывается с использованием синтаксиса instance.method().'),
('Что такое пустой интерфейс interface{} в Go?', '["Интерфейс без методов", "Интерфейс с одним методом", "Интерфейс для работы с каналами", "Интерфейс для работы с файлами"]', 0, 'Пустой интерфейс interface{} не содержит методов и может быть использован для хранения значения любого типа.'),
('Как объявить метод с указателем на структуру в качестве получателя в Go?', '["func (s StructType) MethodName()", "func (s *StructType) MethodName()", "func MethodName(s StructType)", "func MethodName(s *StructType)"]', 1, 'Метод с указателем на структуру в качестве получателя объявляется с использованием синтаксиса func (s *StructType) MethodName().'),
('Как объявить анонимную функцию в Go?', '["func() {}", "() => {}", "function() {}", "func {}"]', 0, 'Анонимная функция объявляется с использованием синтаксиса func() {}.'),
('Как определить метод, который может принимать параметры любого типа в Go?', '["func MethodName[T any](param T)", "func MethodName(param interface{})", "func MethodName(param ...interface{})", "func MethodName(param T)"]', 2, 'Метод, который может принимать параметры любого типа, определяется с использованием синтаксиса func MethodName(param ...interface{}).'),

('Что произойдет, если вызвать метод на nil указателе в Go?', '["Программа завершится с паникой", "Метод будет вызван без ошибок", "Вернется nil", "Будет вызван метод nil объекта"]', 0, 'Если вызвать метод на nil указателе, программа завершится с паникой, если метод не обрабатывает nil указатели явно.'),
('Какой метод из пакета sync.WaitGroup уменьшает счетчик горутин на единицу?', '["Done()", "Add()", "Wait()", "Subtract()"]', 0, 'Метод Done() уменьшает счетчик горутин на единицу в пакете sync.WaitGroup.'),
('Что произойдет, если закрыть канал, и потом попытаться отправить в него значение?', '["Будет отправлено значение", "Программа завершится с паникой", "Ничего не произойдет", "Будет создан новый канал"]', 1, 'Если закрыть канал и потом попытаться отправить в него значение, программа завершится с паникой.'),
('Что произойдет, если не вызвать метод Done() на WaitGroup в Go?', '["Программа завершится с ошибкой", "Программа будет ждать бесконечно", "Ничего не произойдет", "Будет вызвана паника"]', 1, 'Если не вызвать метод Done() на WaitGroup, программа будет ждать бесконечно, так как WaitGroup никогда не достигнет нуля.'),
('Какой метод используется для преобразования строки в число с плавающей точкой?', '["strconv.Atoi()", "strconv.ParseFloat()", "strconv.ParseInt()", "strconv.FormatFloat()"]', 1, 'Метод strconv.ParseFloat() используется для преобразования строки в число с плавающей точкой.'),
('Что делает ключевое слово fallthrough в конструкции switch?', '["Завершает выполнение switch", "Продолжает выполнение следующего case", "Вызывает панику", "Пропускает текущий case"]', 1, 'Ключевое слово fallthrough продолжает выполнение следующего case в конструкции switch.'),
('Как объявить метод, который принимает переменное количество аргументов одного типа?', '["func MethodName(args ...Type)", "func MethodName(args []Type)", "func MethodName(args ...interface{})", "func MethodName(args ...T)"]', 0, 'Метод, который принимает переменное количество аргументов одного типа, объявляется с использованием синтаксиса func MethodName(args ...Type).'),
('Какое значение можно использовать в конструкции select для выполнения по умолчанию?', '["default:", "else:", "case nil:", "switch:"]', 0, 'Значение default: используется в конструкции select для выполнения по умолчанию, если ни один из case не готов к выполнению.'),
('Как объявить метод, который будет доступен только внутри пакета?', '["func methodName() {}", "func MethodName() {}", "private func methodName() {}", "internal func methodName() {}"]', 0, 'Метод, который начинается с маленькой буквы, будет доступен только внутри пакета.'),
('Что произойдет, если вызвать panic внутри defer функции?', '["Ничего не произойдет", "Программа продолжит выполнение", "Другая defer функция будет вызвана", "Программа завершится с паникой"]', 3, 'Если вызвать panic внутри defer функции, программа завершится с паникой после выполнения всех остальных отложенных функций.'),
('Как определить метод, который принимает и возвращает каналы?', '["func MethodName(chan int) chan int", "func MethodName(chan<- int) <-chan int", "func MethodName(chan int) (chan int)", "Все вышеперечисленное"]', 3, 'Все вышеперечисленные способы определяют метод, который принимает и возвращает каналы.'),
('Что делает функция copy() в Go?', '["Копирует данные из строки в строку", "Копирует данные из карты в карту", "Копирует данные из структуры в структуру", "Копирует данные из одного среза в другой"]', 3, 'Функция copy() копирует данные из одного среза в другой.'),

('Что произойдет, если вызвать метод на nil интерфейсе в Go?', '["Программа завершится с паникой", "Метод будет вызван без ошибок", "Вернется nil", "Будет вызван метод nil объекта"]', 0, 'Если вызвать метод на nil интерфейсе, программа завершится с паникой, если метод не обрабатывает nil явно.'),
('Какой тип блокировки обеспечивает одновременное чтение, но эксклюзивное написание?', '["Mutex", "RWMutex", "Once", "WaitGroup"]', 1, 'RWMutex (Read-Write Mutex) обеспечивает одновременное чтение, но эксклюзивное написание.'),
('Что произойдет, если два разных пакета импортируют друг друга?', '["Будет ошибка компиляции", "Программа зависнет", "Будет ошибка времени выполнения", "Программа выполнится без ошибок"]', 0, 'Если два разных пакета импортируют друг друга, будет ошибка компиляции из-за циклической зависимости.'),
('Как объявить канал, который может только отправлять данные?', '["chan int", "<-chan int", "chan<- int", "chan<- int, <-chan int"]', 2, 'Канал, который может только отправлять данные, объявляется как chan<- int.'),
('Какой метод используется для настройки параллелизма выполнения тестов?', '["t.concurrent()", "t.RunParallel()", "t.Run()", "t.Concurrent()"]', 2, 'Метод t.Parallel() используется для настройки параллелизма выполнения тестов.'),
('Что произойдет, если не закрыть канал, который больше не используется?', '["Программа завершится с ошибкой", "Произойдет утечка памяти", "Горутины, ожидающие на этом канале, заблокируются", "Ничего не произойдет"]', 2, 'Если не закрыть канал, горутины, ожидающие на этом канале, заблокируются.'),
('Какой метод из пакета context используется для создания дочернего контекста с тайм-аутом?', '["WithCancel", "WithDeadline", "WithTimeout", "WithValue"]', 2, 'Метод context.WithTimeout используется для создания дочернего контекста с тайм-аутом.'),
('Что произойдет, если вызывать метод Lock() на уже заблокированном мьютексе?', '["Программа завершится с ошибкой", "Произойдет дедлок", "Горутина продолжит выполнение", "Мьютекс автоматически разблокируется"]', 1, 'Если вызвать метод Lock() на уже заблокированном мьютексе, произойдет дедлок, т.к. горутина будет ждать освобождения мьютекса.'),
('Какой тип блокировки обеспечивает однократное выполнение определенного кода?', '["Mutex", "RWMutex", "Once", "WaitGroup"]', 3, 'Once обеспечивает однократное выполнение определенного кода.'),
('Как избежать утечек памяти при использовании таймеров в Go?', '["Использовать defer для остановки таймера", "Использовать таймеры с канала", "Обнулять таймер", "Использовать context для отмены"]', 0, 'Использование defer для остановки таймера помогает избежать утечек памяти.'),
('Какой метод из пакета testing используется для создания под-тестов?', '["Ответы b, d", "SubTest", "Parallel", "Run"]', 0, 'Метод t.Run используется для создания под-тестов в пакете testing.'),
('Что произойдет, если вызвать метод RLock() на RWMutex в Go, когда другой горутин вызвал Lock()?', '["Произойдет дедлок", "Метод RLock() будет блокироваться до освобождения Lock()", "Будет вызвана паника", "Горутина продолжит выполнение"]', 1, 'Метод RLock() будет блокироваться до тех пор, пока Lock() не будет освобожден.'),
('Какой метод используется для завершения работы горутины при закрытии программы?', '["runtime.Goexit()", "runtime.Gosched()", "runtime.GC()", "runtime.NumGoroutine()"]', 0, 'Метод runtime.Goexit() используется для завершения работы текущей горутины.'),
('Какой метод из пакета sync используется для безопасного доступа к целочисленным значениям из нескольких горутин?', '["sync.RWMutex.RLock()", "sync.Mutex.Lock()", "sync.WaitGroup.Wait()", "sync/atomic.AddInt32()"]', 3, 'Метод sync/atomic.AddInt32() используется для безопасного доступа и изменения целочисленных значений из нескольких горутин.'),

('Чем отличается интерфейс от структуры в Go?', '["Интерфейс не содержит методов", "Структура используется только для хранения данных", "Интерфейс определяет только поведение, а структура - данные и поведение", "Структура определяет только поведение"]', 2, 'Интерфейс в Go определяет только поведение, а структура может содержать и данные, и поведение.'),
('Что такое пакетный рантайм (package runtime) в Go?', '["Встроенные функции для работы с пакетами", "Часть стандартной библиотеки, обеспечивающая основные функции исполнения программы", "Интерфейсы для работы с горутинами", "Расширения стандартной библиотеки"]', 1, 'Пакетный рантайм в Go - это часть стандартной библиотеки, которая обеспечивает основные функции исполнения программы, включая управление памятью и горутинами.'),
('Из чего состоит файл go.mod?', '["Список всех используемых пакетов", "Информация о версии Go", "Метаданные проекта и его зависимости", "Исполняемый файл проекта"]', 2, 'Файл go.mod состоит из метаданных проекта и его зависимостей, включая список используемых пакетов и их версии.'),
('Для чего нужен replace в go.mod?', '["Для временного отключения зависимостей", "Для замены стандартных пакетов на пользовательские реализации", "Для перехода на другую версию Go", "Для обновления всех зависимостей"]', 1, 'Ключевое слово replace в go.mod используется для замены стандартных пакетов на пользовательские реализации в процессе разработки или тестирования.'),
('Что означает indirect в go.mod?', '["Зависимость не явно задана в проекте", "Зависимость косвенно используется в проекте", "Зависимость не должна обновляться автоматически", "Зависимость используется только для разработки"]', 1, 'Когда зависимость помечена как indirect в go.mod, это означает, что она косвенно используется в проекте через другие зависимости.'),
('Что такое Big O-нотация?', '["Оценка объема памяти, используемого алгоритмом", "Оценка времени работы алгоритма в среднем случае", "Оценка времени работы алгоритма в худшем случае", "Все вышеперечисленное"]', 2, 'Big O-нотация используется для оценки временной сложности алгоритма в худшем случае.'),
('Какой алгоритм сортировки имеет наилучшую временную сложность в среднем случае среди сравнительных алгоритмов?', '["Пузырьковая сортировка (Bubble Sort)", "Сортировка слиянием (Merge Sort)", "Быстрая сортировка (Quick Sort)", "Вставочная сортировка (Insertion Sort)"]', 2, 'Быстрая сортировка (Quick Sort) имеет наилучшую временную сложность O(n log n) в среднем случае среди сравнительных алгоритмов сортировки.'),

('Что такое gRPC в контексте сетевого взаимодействия в Go?', '["Протокол для передачи SOAP-сообщений", "Формат передачи данных через RESTful API", "Протокол для удаленного вызова процедур (RPC) с использованием HTTP/2", "Фреймворк для разработки веб-приложений"]', 3, 'gRPC - это протокол для удаленного вызова процедур (RPC) с использованием HTTP/2, предназначенный для эффективного взаимодействия между различными сервисами.'),
('Какие преимущества предоставляет gRPC по сравнению с RESTful API?', '["Более высокая производительность и эффективность благодаря использованию HTTP/2 и сериализации Protocol Buffers", "Простота использования и понимания", "Легкость в развертывании", "Все вышеперечисленное"]', 0, 'gRPC предлагает более высокую производительность и эффективность благодаря использованию HTTP/2 и сериализации Protocol Buffers.'),
('Какое преимущество предоставляет использование Protocol Buffers в gRPC?', '["Простота в использовании", "Высокая производительность при передаче данных", "Легкость в развертывании", "Возможность работы с JSON-форматом"]', 1, 'Protocol Buffers предоставляют высокую производительность при передаче данных благодаря компактному бинарному формату.'),
('Какой тип данных в Go часто используется для представления JSON-структур в RESTful API?', '["Структуры", "Массивы", "Карты", "Целые числа"]', 0, 'В Go для представления JSON-структур в RESTful API часто используются структуры данных.'),

('Что включает в себя структура слайса?', '["len, cap", "len и cap", "Ссылку на базовый массив, len и cap", "Слайс это не структура"]', 2, 'Слайс - структура go, которая содержит в себе ссылку на базовый массив, длину (len) и ёмкость (cap).'),
('Чему равны zero-value для слайса?', '["error,", "0, 0", "false", "nil, nil"]', 1, 'Zero-value для слайса - nil, а его длина и ёмкость (len, cap) равны нулю, т.к. "под ним" нет инициализированного масива'),
('Как проверить слайс var a []int на пустоту?', '["fmt.Println(a == nil)", "fmt.Println(!a)", "Стандартная библиотека не даёт проверки слайса на пустоту","fmt.Println(len(a) == 0)"]', 3, 'Самый надежный способ проверить слайс на пустоту - это проверить его длину на ноль.');