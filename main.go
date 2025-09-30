package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"

	. "modernc.org/tk9.0"
)

//go:embed icon.png
var ico []byte

func main() {
	// Animação de 3 pontinhos usando Go Routine e Atomics
	var dots atomic.Value
	dots.Store("")
	go func() {
		seq := []string{"", ".", "..", "..."}
		i := 0
		for {
			dots.Store(seq[i%len(seq)])
			i++
			time.Sleep(300 * time.Millisecond)
		}
	}()

	// Definiçoes da Janela
	App.WmTitle("Ferramenta de BackUp - TI")
	App.SetResizable(false, false)
	WmGeometry(App, "800x600")
	ActivateTheme("clam")

	// Campos compartilhados entre execuções
	l1 := TLabel(Txt("Entre com o usuário Windows (AD):"), Background("#f0f0f0"))
	l1e := TEntry(Textvariable("Digite o usuário.."))
	Grid(l1, Row(0), Column(0), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(l1e, Row(0), Column(1), Sticky("we"), Padx("1m"), Pady("1m"))

	// Opção 01 = Criar um novo backup no servidor
	sep1 := TSeparator(Orient("horizontal"))
	Grid(sep1, Row(1), Column(0), Columnspan(2), Sticky("we"), Padx("0.5m"), Pady("10m"))
	labletitle1 := TLabel(Txt("SUBIR UM NOVO BACKUP PARA O SERVIDOR"), Font("Arial", "15"), Background("#add8e6"))
	Grid(labletitle1, Row(2), Column(0), Columnspan(2), Sticky("ew"), Padx("1m"), Pady("1m"))
	l2 := TLabel(Txt("Escolha o diretório a ser salvo no BackUp:"), Background("#f0f0f0"))
	l2l := TLabel(Textvariable("Aguardando diretório"), Background("#add8e6"))
	StyleConfigure("blue.TButton",
		Background("#b2d2dd"),
		Foreground("#003446"),
		Font("Arial Rounded", "10", "bold"),
	)
	l2b := TButton(Txt("Selecionar Diretório"), Style("blue.TButton"), Command(
		func() {
			dir := ChooseDirectory(Title("Diretório para criar BackUp"))
			l2l.Configure(Textvariable(dir))
		},
	))
	StyleConfigure("Blue.Horizontal.TProgressbar",
		Background("#42798b"),
		Troughcolor(White),
		Bordercolor("#42798b"),
		Lightcolor("#42798b"),
		Darkcolor("#003446"),
	)
	barCriar := TProgressbar(Mode("determinate"), Maximum(5), Style("Blue.Horizontal.TProgressbar"))
	statusCriar := TLabel(Textvariable("OK"))
	chCriar := make(chan []string)
	var criarBKP *TButtonWidget
	criarBKP = TButton(Txt("CRIAR BACKUP"), Style("blue.TButton"), Command(func() {
		criarBKP.Configure(State("disabled"))
		seguir := MessageBox(Title("Confirmar Inicio do Processo"),
			Msg("Processeguir com o processamento dos dados?"),
			Type("yesno"),
			Icon("question"),
		)
		if seguir == "yes" {
			go Simular(chCriar, l1e.Textvariable(), l2l.Textvariable())
		} else {
			criarBKP.Configure(State("enabled"))
		}
	}),
	)
	var data string
	var animate bool
	NewTicker(300*time.Millisecond, func() {
		select {
		case v := <-chCriar:
			data = v[0]
			if data == "OK" {
				criarBKP.Configure(State("enabled"))
				animate = false
				break
			}
			if strings.HasPrefix(data, "ERRO") {
				statusCriar.Configure(Textvariable(data))
				criarBKP.Configure(State("enabled"))
				animate = false
				break
			}
			barCriar.Configure(Value(v[1]))
			animate = true
		default:
			if animate {
				statusCriar.Configure(Textvariable(data + (dots.Load().(string))))
			} else {
				statusCriar.Configure(Textvariable(data))
			}
		}
	})
	Grid(l2, Row(3), Column(0), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(l2b, Row(3), Column(1), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(l2l, Row(4), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(criarBKP, Row(5), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(barCriar, Row(6), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(statusCriar, Row(7), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))

	// Opção 02 = Restaurar um backup do servidor
	sep2 := TSeparator(Orient("horizontal"))
	Grid(sep2, Row(8), Column(0), Columnspan(2), Sticky("we"), Padx("0.5m"), Pady("10m"))
	labletitle2 := TLabel(Txt("RESTAURAR UM BACKUP DO SERVIDOR"), Font("Arial", "15"), Background("#ade6b9"))
	Grid(labletitle2, Row(9), Column(0), Columnspan(2), Sticky("ew"), Padx("1m"), Pady("1m"))
	l3 := TLabel(Txt("Escolha um diretório para restaurar o BackUp:"), Background("#f0f0f0"))
	l3l := TLabel(Textvariable("Aguardando diretório"), Background("#ade6b9"))
	StyleConfigure("green.TButton",
		Background("#c8e6c8"),
		Foreground("#004200"),
		Font("Arial Rounded", "10", "bold"),
	)
	l3b := TButton(Txt("Selecionar Diretório"), Style("green.TButton"), Command(
		func() {
			dir := ChooseDirectory(Title("Diretório para criar BackUp"))
			l3l.Configure(Textvariable(dir))
		},
	))
	StyleConfigure("Green.Horizontal.TProgressbar",
		Background("#458d45"),
		Troughcolor(White),
		Bordercolor("#458d45"),
		Lightcolor("#458d45"),
		Darkcolor("#004200"),
	)
	barRestaurar := TProgressbar(Orient("horizontal"), Mode("determinate"), Maximum(5), Style("Green.Horizontal.TProgressbar"))
	statusRestaurar := TLabel(Textvariable("OK"))
	chRestaurar := make(chan []string)
	var restaurarBKP *TButtonWidget
	restaurarBKP = TButton(Txt("RESTAURAR BACKUP"), Style("green.TButton"), Command(func() {
		restaurarBKP.Configure(State("disabled"))
		seguir := MessageBox(Title("Confirmar Inicio do Processo"),
			Msg("Processeguir com o processamento dos dados?"),
			Type("yesno"),
			Icon("question"),
		)
		if seguir == "yes" {
			go Simular(chRestaurar, l1e.Textvariable(), l3l.Textvariable())
		} else {
			restaurarBKP.Configure(State("enabled"))
		}
	}))
	var data1 string
	var animate1 bool
	NewTicker(100*time.Millisecond, func() {
		select {
		case v := <-chRestaurar:
			data1 = v[0]
			if data1 == "OK" {
				restaurarBKP.Configure(State("enabled"))
				animate1 = false
				break
			}
			if strings.HasPrefix(data1, "ERRO") {
				statusRestaurar.Configure(Textvariable(data1))
				restaurarBKP.Configure(State("enabled"))
				animate1 = false
				break
			}
			barRestaurar.Configure(Value(v[1]))
			animate1 = true
		default:
			if animate1 {
				statusRestaurar.Configure(Textvariable(data1 + (dots.Load().(string))))
			} else {
				statusRestaurar.Configure(Textvariable(data1))
			}
		}
	})
	Grid(l3, Row(10), Column(0), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(l3b, Row(10), Column(1), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(l3l, Row(11), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(restaurarBKP, Row(12), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(barRestaurar, Row(13), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))
	Grid(statusRestaurar, Row(14), Column(0), Columnspan(2), Sticky("we"), Padx("1m"), Pady("1m"))

	GridColumnConfigure(App, 1, Weight(1))

	App.IconPhoto(NewPhoto(Data(ico)))
	App.Center().Wait()
}

func Simular(ch chan []string, user string, path string) {
	ch <- []string{fmt.Sprintf("Iniciando o Processo do usuario %s", user), "0"}
	if _, err := os.Stat(path); os.IsNotExist(err) || path == "Aguardando diretório" {
		ch <- []string{"ERRO: O diretório informado não existe.", "0"}
		return
	}
	if user == "" || user == "Digite o usuário.." {
		ch <- []string{"ERRO: Informe um usuário válido.", "0"}
		return
	}
	time.Sleep(5 * time.Second)
	ch <- []string{"Etapa 01 - Verificando arquivos", "1"}
	time.Sleep(5 * time.Second)
	ch <- []string{"Etapa 02 - Compactando dados", "2"}
	time.Sleep(5 * time.Second)
	ch <- []string{"Etapa 03 - Conectando ao servidor", "3"}
	time.Sleep(5 * time.Second)
	spl := strings.Split(path, "/")
	ch <- []string{fmt.Sprintf("Etapa 04 - Processando o diretório: %s", spl[len(spl)-1]), "4"}
	time.Sleep(5 * time.Second)
	ch <- []string{"OK", "5"}
}
