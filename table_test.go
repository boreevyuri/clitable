package clitable

import "testing"

func TestSimpleHeader(t *testing.T) {
	table := NewTable("id")
	header :=
		"+--+\n" +
			"|id|\n" +
			"+--+\n"

	t.Log(table.String())
	t.Log(header)
	if table.String() != header {
		t.Fail()
	}
}

func TestSimpleHeaderWithPadding(t *testing.T) {
	table := NewTable("id")
	column := table.GetColumnByName("id")
	column.HeaderStyle = &ColumnStyle{
		PaddingLeft:  2,
		PaddingRight: 1,
	}

	header :=
		"+-----+\n" +
			"|  id |\n" +
			"+-----+\n"

	t.Log(table.String())
	t.Log(header)
	if table.String() != header {
		t.Fail()
	}
}

func TestHeadersWithPadding(t *testing.T) {
	table := NewTable("id", "name")
	column := table.GetColumnByName("id")
	column.HeaderStyle = &ColumnStyle{
		PaddingLeft:  2,
		PaddingRight: 1,
	}

	header :=
		"+-----+----+\n" +
			"|  id |name|\n" +
			"+-----+----+\n"

	t.Log(table.String())
	t.Log(header)
	if table.String() != header {
		t.Fail()
	}
}

func TestMultiLineHeader(t *testing.T) {
	WinSize.Col = 24
	table := NewTable("id", "name", "too long header super name")
	column := table.GetColumnByName("id")
	column.HeaderStyle = &ColumnStyle{
		PaddingLeft:  2,
		PaddingRight: 1,
	}

	header :=
		"+-----+----+------+\n" +
			"|  id |    | too  |\n" +
			"|     |    | long |\n" +
			"|     |name|header|\n" +
			"|     |    |super |\n" +
			"|     |    | name |\n" +
			"+-----+----+------+\n"

	tableStr := table.String()
	t.Log(tableStr)
	t.Log(header)
	if tableStr != header {
		t.Fail()
	}
}

func TestMultiLineHeaderWithBody(t *testing.T) {
	WinSize.Col = 31
	table := NewTable("id", "name", "too long header super name")
	column := table.GetColumnByName("id")
	column.HeaderStyle = &ColumnStyle{
		PaddingLeft:  2,
		PaddingRight: 1,
	}
	column = table.GetColumnByName("name")
	column.BodyStyle = &ColumnStyle{
		Align: ColumnAlignRight,
	}
	column = table.GetColumnByName("too long header super name")
	column.BodyStyle = &ColumnStyle{
		Align: ColumnAlignCenter,
	}

	table.AddRow(1, "firstname", "-")
	table.AddRow(2, "lastname", "--")
	table.AddRow(123, "address", "---")
	table.AddRow(1234567, "city", "----")

	header :=
		"+-------+---------+------+\n" +
			"|  id   |         | too  |\n" +
			"|       |         | long |\n" +
			"|       |  name   |header|\n" +
			"|       |         |super |\n" +
			"|       |         | name |\n" +
			"+-------+---------+------+\n" +
			"|1      |firstname|  -   |\n" +
			"+-------+---------+------+\n" +
			"|2      | lastname|  --  |\n" +
			"+-------+---------+------+\n" +
			"|123    |  address| ---  |\n" +
			"+-------+---------+------+\n" +
			"|1234567|     city| ---- |\n" +
			"+-------+---------+------+\n"

	tableStr := table.String()
	t.Log(tableStr)
	t.Log(header)
	if tableStr != header {
		t.Fail()
	}
}

func TestMultiLineBody(t *testing.T) {
	WinSize.Col = 105
	table := NewTable("id", "name", "description", "short description")
	column := table.GetColumnByName("id")
	column.HeaderStyle = &ColumnStyle{
		PaddingLeft:  2,
		PaddingRight: 1,
	}
	column = table.GetColumnByName("name")
	column.BodyStyle = &ColumnStyle{
		Align:        ColumnAlignRight,
		PaddingLeft:  1,
		PaddingRight: 1,
	}
	column = table.GetColumnByName("short description")
	column.BodyStyle = &ColumnStyle{
		Align:         ColumnAlignCenter,
		VerticalAlign: ColumnVerticalAlignBottom,
		PaddingTop:    2,
		PaddingBottom: 1,
	}

	table.AddRow(
		1,
		"DecodeLastRune",
		"DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it returns (RuneError, 1). Both are impossible results for correct UTF-8."+
			"An encoding is invalid if it is incorrect UTF-8, encodes a rune that is out of range, or is not the shortest possible UTF-8 encoding for the value. No other validation is performed.."+
			"DecodeLastRuneInString is like DecodeLastRune but its input is a string. If s is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it returns (RuneError, 1). Both are impossible results for correct UTF-8."+
			"An encoding is invalid if it is incorrect UTF-8, encodes a rune that is out of range, or is not the shortest possible UTF-8 encoding for the value.",
		"DecodeRune unpacks the first UTF-8 encoding in p and returns the rune and its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it returns (RuneError, 1).",
	)
	table.AddRow(
		2,
		"DecodeLastRuneInString",
		"DecodeLastRune unpacks the last UTF-8 encoding in p and returns the rune and its width in bytes. If p is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it returns (RuneError, 1). Both are impossible results for correct UTF-8.",
		"--",
	)

	header :=
		"+-----+------------------------+------------------------------------------------+-------------------+\n" +
			"|  id |          name          |                  description                   | short description |\n" +
			"+-----+------------------------+------------------------------------------------+-------------------+\n" +
			"|1    |         DecodeLastRune |DecodeLastRune unpacks the last UTF-8 encoding  |                   |\n" +
			"|     |                        |in p and returns the rune and its width in      |                   |\n" +
			"|     |                        |bytes. If p is empty it returns (RuneError, 0). |                   |\n" +
			"|     |                        |Otherwise, if the encoding is invalid, it       |                   |\n" +
			"|     |                        |returns (RuneError, 1). Both are impossible     |                   |\n" +
			"|     |                        |results for correct UTF-8.An encoding is        |                   |\n" +
			"|     |                        |invalid if it is incorrect UTF-8, encodes a     |DecodeRune unpacks |\n" +
			"|     |                        |rune that is out of range, or is not the        |  the first UTF-8  |\n" +
			"|     |                        |shortest possible UTF-8 encoding for the value. | encoding in p and |\n" +
			"|     |                        |No other validation is                          | returns the rune  |\n" +
			"|     |                        |performed..DecodeLastRuneInString is like       | and its width in  |\n" +
			"|     |                        |DecodeLastRune but its input is a string. If s  |  bytes. If p is   |\n" +
			"|     |                        |is empty it returns (RuneError, 0). Otherwise,  | empty it returns  |\n" +
			"|     |                        |if the encoding is invalid, it returns          |  (RuneError, 0).  |\n" +
			"|     |                        |(RuneError, 1). Both are impossible results for | Otherwise, if the |\n" +
			"|     |                        |correct UTF-8.An encoding is invalid if it is   |    encoding is    |\n" +
			"|     |                        |incorrect UTF-8, encodes a rune that is out of  |    invalid, it    |\n" +
			"|     |                        |range, or is not the shortest possible UTF-8    |      returns      |\n" +
			"|     |                        |encoding for the value.                         |  (RuneError, 1).  |\n" +
			"+-----+------------------------+------------------------------------------------+-------------------+\n" +
			"|2    | DecodeLastRuneInString |DecodeLastRune unpacks the last UTF-8 encoding  |                   |\n" +
			"|     |                        |in p and returns the rune and its width in      |                   |\n" +
			"|     |                        |bytes. If p is empty it returns (RuneError, 0). |                   |\n" +
			"|     |                        |Otherwise, if the encoding is invalid, it       |                   |\n" +
			"|     |                        |returns (RuneError, 1). Both are impossible     |        --         |\n" +
			"|     |                        |results for correct UTF-8.                      |                   |\n" +
			"+-----+------------------------+------------------------------------------------+-------------------+\n"

	tableStr := table.String()
	t.Log(tableStr)
	t.Log(header)
	if tableStr != header {
		t.Fail()
	}
}

func TestCyrillicMultiLineBody(t *testing.T) {
	WinSize.Col = 155
	table := NewTable("#", "Имя", "Описание", "Короткое описание")
	column := table.GetColumnByName("#")
	column.HeaderStyle = &ColumnStyle{
		PaddingLeft:  2,
		PaddingRight: 1,
	}
	column = table.GetColumnByName("Имя")
	column.BodyStyle = &ColumnStyle{
		Align:        ColumnAlignRight,
		PaddingLeft:  1,
		PaddingRight: 1,
	}
	column = table.GetColumnByName("Короткое описание")
	column.BodyStyle = &ColumnStyle{
		Align:         ColumnAlignCenter,
		VerticalAlign: ColumnVerticalAlignBottom,
		PaddingTop:    2,
		PaddingBottom: 1,
	}

	table.AddRow(
		1,
		"Golang",
		"Go (часто также Golang) — компилируемый, многопоточный язык программирования, разработанный компанией Google[2]. Первоначальная разработка Go началась в сентябре 2007 года, а его непосредственным проектированием занимались Роберт Гризмер, Роб Пайк и Кен Томпсон[3] занимавшиеся до этого проектом разработки операционной системы Inferno. Официально язык был представлен в ноябре 2009 года. На данный момент его поддержка осуществляется для операционных систем: FreeBSD, OpenBSD, Linux, Mac OS X, Windows[4], начиная с версии 1.3 в язык Go включена экспериментальная поддержка DragonFly BSD, Plan 9 и Solaris, начиная с версии 1.4 поддержка платформы Android.",
		"Следует отметить, что название языка, выбранное компанией Google, практически совпадает с названием языка программирования Go!, созданного Ф. Джи. МакКейбом и К. Л. Кларком в 2003 году.[5] Обсуждение названия ведётся на странице, посвящённой Go[5].",
	)
	table.AddRow(
		2,
		"Java",
		"Java[11] — объектно-ориентированный язык программирования, разработанный компанией Sun Microsystems (в последующем приобретённой компанией Oracle). Приложения Java обычно транслируются в специальный байт-код, поэтому они могут работать на любой виртуальной Java-машине вне зависимости от компьютерной архитектуры. Дата официального выпуска — 23 мая 1995 года. Изначально язык назывался Oak («Дуб») разрабатывался Джеймсом Гослингом для программирования бытовых электронных устройств. Впоследствии он был переименован в Java и стал использоваться для написания клиентских приложений и серверного программного обеспечения. Назван в честь марки кофе Java, которая, в свою очередь, получила наименование одноимённого острова (Ява), поэтому на официальной эмблеме языка изображена чашка с парящим кофе. Существует и другая версия происхождения названия языка, связанная с аллюзией на кофе-машину как пример бытового устройства, для программирования которого изначально язык создавался.",
		"Программы на Java транслируются в байт-код, выполняемый виртуальной машиной Java (JVM) — программой, обрабатывающей байтовый код и передающей инструкции оборудованию как интерпретатор.",
	)
	table.AddRow(
		3,
		"PHP",
		"PHP (англ. PHP: Hypertext Preprocessor — «PHP: препроцессор гипертекста»; первоначально Personal Home Page Tools[4] — «Инструменты для создания персональных веб-страниц»; произносится пи-эйч-пи) — скриптовый язык[5] программирования общего назначения, интенсивно применяемый для разработки веб-приложений. В настоящее время поддерживается подавляющим большинством хостинг-провайдеров и является одним из лидеров среди языков программирования, применяющихся для создания динамических веб-сайтов[6].Язык и его интерпретатор разрабатываются группой энтузиастов в рамках проекта с открытым кодом[7]. Проект распространяется под собственной лицензией, несовместимой с GNU GPL.",
		"Синтаксис PHP подобен синтаксису языка Си. Некоторые элементы, такие как ассоциативные массивы и цикл foreach, заимствованы из Perl. Для работы программы не требуется описывать какие-либо переменные, используемые модули и т. п. Любая программа может начинаться непосредственно с оператора PHP.",
	)

	header :=
		"+----+--------+----------------------------------------------------------------------------------------------------+----------------------------------+\n" +
			"|  # |  Имя   |                                              Описание                                              |        Короткое описание         |\n" +
			"+----+--------+----------------------------------------------------------------------------------------------------+----------------------------------+\n" +
			"|1   | Golang |Go (часто также Golang) — компилируемый, многопоточный язык программирования, разработанный         |                                  |\n" +
			"|    |        |компанией Google[2]. Первоначальная разработка Go началась в сентябре 2007 года, а его              |                                  |\n" +
			"|    |        |непосредственным проектированием занимались Роберт Гризмер, Роб Пайк и Кен Томпсон[3] занимавшиеся  |                                  |\n" +
			"|    |        |до этого проектом разработки операционной системы Inferno. Официально язык был представлен в ноябре |  Следует отметить, что название  |\n" +
			"|    |        |2009 года. На данный момент его поддержка осуществляется для операционных систем: FreeBSD, OpenBSD, |    языка, выбранное компанией    |\n" +
			"|    |        |Linux, Mac OS X, Windows[4], начиная с версии 1.3 в язык Go включена экспериментальная поддержка    | Google, практически совпадает с  |\n" +
			"|    |        |DragonFly BSD, Plan 9 и Solaris, начиная с версии 1.4 поддержка платформы Android.                  | названием языка программирования |\n" +
			"|    |        |                                                                                                    |Go!, созданного Ф. Джи. МакКейбом |\n" +
			"|    |        |                                                                                                    | и К. Л. Кларком в 2003 году.[5]  |\n" +
			"|    |        |                                                                                                    |  Обсуждение названия ведётся на  |\n" +
			"|    |        |                                                                                                    |   странице, посвящённой Go[5].   |\n" +
			"+----+--------+----------------------------------------------------------------------------------------------------+----------------------------------+\n" +
			"|2   |   Java |Java[11] — объектно-ориентированный язык программирования, разработанный компанией Sun Microsystems |                                  |\n" +
			"|    |        |(в последующем приобретённой компанией Oracle). Приложения Java обычно транслируются в специальный  |                                  |\n" +
			"|    |        |байт-код, поэтому они могут работать на любой виртуальной Java-машине вне зависимости от            |                                  |\n" +
			"|    |        |компьютерной архитектуры. Дата официального выпуска — 23 мая 1995 года. Изначально язык назывался   |                                  |\n" +
			"|    |        |Oak («Дуб») разрабатывался Джеймсом Гослингом для программирования бытовых электронных устройств.   |                                  |\n" +
			"|    |        |Впоследствии он был переименован в Java и стал использоваться для написания клиентских приложений и |Программы на Java транслируются в |\n" +
			"|    |        |серверного программного обеспечения. Назван в честь марки кофе Java, которая, в свою очередь,       |байт-код, выполняемый виртуальной |\n" +
			"|    |        |получила наименование одноимённого острова (Ява), поэтому на официальной эмблеме языка изображена   | машиной Java (JVM) — программой, |\n" +
			"|    |        |чашка с парящим кофе. Существует и другая версия происхождения названия языка, связанная с аллюзией |  обрабатывающей байтовый код и   |\n" +
			"|    |        |на кофе-машину как пример бытового устройства, для программирования которого изначально язык        |      передающей инструкции       |\n" +
			"|    |        |создавался.                                                                                         | оборудованию как интерпретатор.  |\n" +
			"+----+--------+----------------------------------------------------------------------------------------------------+----------------------------------+\n" +
			"|3   |    PHP |PHP (англ. PHP: Hypertext Preprocessor — «PHP: препроцессор гипертекста»; первоначально Personal    |                                  |\n" +
			"|    |        |Home Page Tools[4] — «Инструменты для создания персональных веб-страниц»; произносится пи-эйч-пи) — |                                  |\n" +
			"|    |        |скриптовый язык[5] программирования общего назначения, интенсивно применяемый для разработки        |                                  |\n" +
			"|    |        |веб-приложений. В настоящее время поддерживается подавляющим большинством хостинг-провайдеров и     | Синтаксис PHP подобен синтаксису |\n" +
			"|    |        |является одним из лидеров среди языков программирования, применяющихся для создания динамических    |  языка Си. Некоторые элементы,   |\n" +
			"|    |        |веб-сайтов[6].Язык и его интерпретатор разрабатываются группой энтузиастов в рамках проекта с       |такие как ассоциативные массивы и |\n" +
			"|    |        |открытым кодом[7]. Проект распространяется под собственной лицензией, несовместимой с GNU GPL.      |  цикл foreach, заимствованы из   |\n" +
			"|    |        |                                                                                                    |  Perl. Для работы программы не   |\n" +
			"|    |        |                                                                                                    |  требуется описывать какие-либо  |\n" +
			"|    |        |                                                                                                    |переменные, используемые модули и |\n" +
			"|    |        |                                                                                                    |   т. п. Любая программа может    |\n" +
			"|    |        |                                                                                                    |   начинаться непосредственно с   |\n" +
			"|    |        |                                                                                                    |          оператора PHP.          |\n" +
			"+----+--------+----------------------------------------------------------------------------------------------------+----------------------------------+\n"

	tableStr := table.String()
	t.Log(tableStr)
	t.Log(header)
	if tableStr != header {
		t.Fail()
	}
}

func TestCustomBorderAndCorner(t *testing.T) {
	table := NewTable("id")
	table.Style = &TableStyle{
		Corner:           "*",
		VerticalBorder:   "!",
		HorizontalBorder: "~",
	}
	column := table.GetColumnByName("id")
	column.HeaderStyle = &ColumnStyle{
		PaddingLeft:  2,
		PaddingRight: 1,
	}

	header :=
		"*~~~~~*\n" +
			"!  id !\n" +
			"*~~~~~*\n"

	t.Log(table.String())
	t.Log(header)
	if table.String() != header {
		t.Fail()
	}
}
