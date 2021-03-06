package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/tophatsteve/saas/service"
	"google.golang.org/grpc"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	lis, err := net.Listen("tcp", ":9191")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	saas.RegisterSausageSvcServer(grpcServer,
		saas.NewService(getSausageList()),
	)

	go func() {
		grpcServer.Serve(lis)
	}()

	<-stop
	grpcServer.GracefulStop()
}

func getSausageList() []string {
	sausages := []string{
		"Ahle Wurst",
		"Alheira",
		"Andouille",
		"Andouillette",
		"Androlla",
		"Azaruja sausage",
		"Balkenbrij",
		"Banat",
		"Beutelwurst",
		"Bierschinken",
		"Bierwurst",
		"Biroldo",
		"Black pudding",
		"Bloedworst",
		"Blunze",
		"Blutwurst",
		"Bockwurst",
		"Boerewors",
		"Botillo",
		"Boudin blanc de Rethel",
		"Boudin",
		"Braadworst",
		"Bratwurst",
		"Braughing",
		"Braunschweiger",
		"Bregenwurst",
		"Brühwurst",
		"Butifarra Soledeñas",
		"Butifarra",
		"Cabanossi",
		"Cervelas de Lyon",
		"Cervelat",
		"Cervelatwurst",
		"Češnovka",
		"Chipolata",
		"Chistorra",
		"Chorizo de Pamplona",
		"Chorizo",
		"Chouriço doce",
		"Chouriço",
		"Ciauscolo",
		"Ciavàr",
		"Cotechino Modena",
		"Cotechino",
		"Csabai",
		"Cumberland",
		"Currywurst",
		"Debrecener",
		"Devon",
		"Diot",
		"Doktorskaya kolbasa",
		"Drisheen",
		"Droëwors",
		"Embutido",
		"Extrawurst",
		"Falukorv",
		"Farinheira",
		"Fläskkorv",
		"Frankfurter Rindswurst",
		"Frankfurter Würstchen",
		"Frikandel",
		"Fuet",
		"Gelbwurst",
		"Genoa salami",
		"Gloucester",
		"Gyulai",
		"Ham sausage",
		"Hog's pudding",
		"Hungarian sausages",
		"Isterband",
		"Jagdwurst",
		"Kabanos",
		"Kaminwurz",
		"Käsekrainer",
		"Kaszanka",
		"Kazy",
		"Kielbasa",
		"Kishka",
		"Knackwurst",
		"Knipp",
		"Kochwurst",
		"Kohlwurst",
		"Krakowska",
		"Krestyanskaya kolbasa",
		"Krupniok",
		"Kulen",
		"Kupati",
		"Landjäger",
		"Lao sausage",
		"Leberwurst",
		"Likëngë",
		"Lincolnshire",
		"Linguiça",
		"Liverwurst",
		"Longaniza",
		"Lukanka",
		"Lunenburg pudding",
		"Makanek",
		"Manchester",
		"Marylebone",
		"Medisterpølse",
		"Merguez",
		"Mettwurst",
		"Metworst",
		"Morcilla",
		"Morcón",
		"Moronga",
		"Mortadella",
		"Morteau sausage",
		"Mustamakkara",
		"Myśliwska",
		"Nădlac",
		"Naem",
		"Nduja",
		"Newmarket",
		"Noumboulo",
		"Nürnberger Bratwürste",
		"Ossenworst",
		"Oxford",
		"Paio",
		"Pick",
		"Pinkel",
		"Potatiskorv",
		"Prasky",
		"Prinskorv",
		"Regensburger Wurst",
		"Rød pølse",
		"Rookworst",
		"Rosette de Lyon",
		"Ryynimakkara",
		"Sabodet",
		"Sai krok Isan",
		"Sai ua",
		"Salami",
		"Salamin",
		"Salchicha",
		"Salchichón",
		"Saucisse de choux",
		"Saucisse de Toulouse",
		"Saucisson de Lyon",
		"Saucisson Vaudois",
		"Saucisson",
		"Saumagen",
		"Saveloy",
		"Schüblig",
		"Sibiu",
		"Siskonmakkara",
		"Skilandis",
		"Snorkers",
		"Sobrasada",
		"Som moo",
		"Soppressata",
		"Sopressa",
		"Sremska kobasica",
		"St. Galler bratwurst",
		"Stippgrütze",
		"Stonner kebab",
		"Strolghino",
		"Sujuk",
		"Sundae",
		"Švargl",
		"Teewurst",
		"Thüringer rotwurst",
		"Tobă",
		"Verivorst",
		"Vienna sausage",
		"Weckewerk",
		"Weisswurst",
		"Westfälische Rinderwurst",
		"White pudding",
		"Wienerwurst",
		"Wienerwürstchen",
		"Winter salami",
		"Wollwurst",
		"Yorkshire",
		"Zungenwurst",
		"Zwiebelwurst",
	}

	return sausages
}
