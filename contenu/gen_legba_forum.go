package main

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/yaml.v3"
)

type Article struct {
	Title      string
	Authors    string
	DatePost   string
	AuthorPost string
	Abstract   string
}
type Post struct {
	Server  string
	Group   string
	Date    time.Time
	Author  string
	Subject string
	Content string
}

// FIXME TODO check login, addServer and group
// FIXME TODO check dates
var loginMath = "amath"
var loginKuip = "jkuip"
var loginHers = "ehers"

// pour les articles ou AuhtorPost est ""
var logDivers = []string{"dyuon", "nroug", "vpoco", "uandr"}
var addrServer = "archive.legba.d22.eu"
var groupMathison = ""

// les autre articles de artMaahison devront être postés avant
var lastDateMathison = "2000-06-30T14:24:01"

// les articles divers sont postés avant
var lastDateOther = "2000-07-15T18:00:00"

var artMathison = []Article{
	{"Transcendance sous co-routines.",
		"D. Yuong, E. Herswing et A. Mathison.",
		"",
		loginMath,
		`Des expériences préliminaires sur la Transcendance Numérique (TN) permettent
la formulation d'une nouvelle conjoncture. Lors des phases terminales
d'auto-corections des caractéristiques de personnalité, il paraît
primordial de passer par des co-routines évolutionnaires de Spinksy-Yuong
avec des timeout ré-évalués par adaptation prodonde. Nos premiers résultats
montrent en effet que sans l'interlaçage de ces co-routines, un effondrement
neuro-dépressif acquiert une vraissemblance voisinant les 92.34 pourcents.
Notons que nos expériences ont été réalisées in vitro, sur le cluster de
calcul SaberSapience v 3.0.23.`},
	{"Transcendance et effondrement neuro-psychologique.",

		"A. Mathison.",
		"",
		loginMath,

		`Nous présentons une review de la littérature récentes sur la
«dégénérescence accélérée» issues de la théorie formulée par le laboratoire
@HigherMind(sous mécénat de Legba Voodoocom). Ces travaux, concommitants à
ceux réalisées par notre propre équipe de SiliconSpirit, ont l'avantage
d'avoir défrichés des pistes non-viables. Il ressort en effet de notre
analyse que la complexité des algorithmes de compressions des aires
thalamiques n'est pas un élément essentiel sur la voie de l'élévation de la
conscience numérique. Cela renforce notre hypothèse selon laquelle l'alignement
mémoriel de classe IV est une contrainte incontournable malgré le surcoût
computationnel certain.`},

	{"Deux encodages mémoriels pour les aires thalamiques intérmédiaires.",

		"Y. Levain, A.P. Revertin, et J.O. Galakievicz.",
		"",
		loginMath,

		`Nous présentons deux algorithmes en O( n2+log(g) ), probablement
epsilon-corrects, pour l'encodage numérique in vivo des activités mémorielles
et sous-conscentes des aires thalammiques humaines du cerveau humain. Le premier
s'appuie sur la librairie sous licence de Gantrell-HypeX, (récemment acquise par
Legba Voodoocom). Le second, dont l'espérence en fiabilité est légèrement
inférieure, sauf pour certains sous-types d'aires thalamiques, est entièrement
nouveau, et libre de droit de propriété.`},
	{"Un bootstrap efficace de l'ontologie phénoménologique.",

		"J. Kuipers et A. Trebinsky.",
		"",
		loginKuip,

		`L'un des écueils dans la Transcendance Numérique Forte est en passe
d'être levé. Nos premiers résultats expérimentaux, en simulation, montrent en
effet que notre méthode de bootstrap de l'ontologie phénoménologique primaire
permet une croissance quasi-exponentielle de la concordance proto-symbolique.
Il en découle logiquement une voie ouverte et prometteuse pour des
extraction-compressions réussie de la conscience humaine. Soulignons que cette
méta-accroissement peut se réaliser sans alignemnt mémoriel bas-niveaux, un
processus dont la compléxité doublement-exponentielle rend l'application pour
le moins délicate.`},
	{"Une voie nouvelle pour la Conscience Numérique : la trans-fusion.",

		"L. Saint-Janvier, D.L. Yu, C. Wu, O. Wellit-Ashley et L. Modina.",
		"",
		loginKuip,

		`A rebours des recherches "mainstream" dans le domaine de la Transcendance
Numérique (TN), nous présentons un cadre formel pour une voie originale et sobre
vers une Conscience Numérique Forte (au sens de Lashley). Notre idée, formulée
dans la logique épistémique modale de second ordre, s'appuie sur la fusion
multi-modale d'une IA de type A et d'une proto-numérisation des couches
superficielles et intermédiaire d'une personnalité humain handi-adaptée. Nous
avons passé notre proposition dans tous les vérificateurs symboliques de la base
de Lashley et tous donnent une probabilité de réussite dépassant les 60%.`},
	{"Du déficit phénoménologique inéluctable des IA de classe A.",

		"R.D. San-Jorgeu, J. Kuipers et A. Trebinsky.",
		"",
		loginKuip,

		`Nos travaux sur le déficit phénoménologique des IA de classe A montrent
la faiblesse des approches trans-humain et handi-humaines. Nous argumentons
notamment sur les dangers, éthiques et moraux, des travaux mêlant handi-adaption
et ancrage artificielle pour l'émergence de conscience. Au delà de considérations
éthiques, nous mettons en exergue une analyse réfutatoire, par la méthode des
quanta-qbits transitionnels, qui démontre l'imfaisabilité des travaux de
[Saint-Janvier et. al] sur cette trans-fusion.`},
	{"Symposium Inter-Coprporatiste sur la Transcendance Numérique.",

		"A. Pernu, C.H. Itchinson et P. Kanakuna.",
		"",
		loginMath,

		`Ces deux journées de dialogues sur les dernières avancées en matie de
Transcendance Numériques furent profitables et porteuses de nouvelles voies de
recherche. Nous avons eu l'honneur d'animer ces journées qui permis des
propositions originales et pertinentes. Citons notamment les bootstrap
phénoménologiques de [Juipers et. al.], la trans-fusion IA/handihumanité
[Saint-Janvier et al.] et l'alignement thalamique de [A. Mathison].
Ce symposiuma eu lieu sous le patronnage de Legba Voodoocom et de SiliconSpirit.`},
	{"Fermeture du Projet Mandrake.",
		"J. Kuipers.",
		lastDateMathison,
		loginKuip,

		`Après une analyse des travaux du Projet «Mandrake», j'ai décidé de le
clore et de transferer son budget au Projet «Phénomos». Le Dr. Alan Mathison
est relevé de son poste de directeur et ses accès révoqués. A toutes fin utiles,
j'ai demandé l'archivages du Projet «Mandrake».`},
}

var artDivers = []Article{
	{"Identification automatique des sources des notices zoologiques du Speculum naturale de Vincent de Beauvais ",
		"Étienne Cuvelier ; Sébastien de Valeriola ; Céline Engelbeen",
		"",
		"",
		`Avec son encyclopédie intitulée Speculum maius, le dominicain du xiiie
siècle Vincent de Beauvais tente de constituer une synthèse générale
 du savoir. Pour ce faire, il rassemble des renseignements provenant
 d’une multitude de sources différentes, chrétiennes et païennes,
 antiques et médiévales. La plupart des notices de son œuvre contiennent
 une mention explicite des sources dont elles sont inspirées, à la
différence de beaucoup des encyclopédies médiévales. Cette caractéristique
permet d’utiliser le Speculum maius comme base d’expérimentation, et de
 lui appliquer des techniques d’apprentissage supervisé et de fouille de
 textes dans le but de relier automatiquement les notices encyclopédiques à
leurs sources. Dans cet article, nous nous livrons à cet exercice pour les
 livres zoologiques de cette encyclopédie et nous analysons ensuite les apports,
 les limites et les perspectives des résultats obtenus dans l’optique d’une
 application future à d’autres encyclopédies dont les notices ne mentionnent
 pas leurs sources.`},
	{"Modéliser ce qui résiste à la modélisation",
		"A. Bénel",
		"",
		"",
		`Pour traiter de l’intelligence artificielle appliquée aux humanités numériques,
 cet article prend le pari risqué de concentrer son état de l’art sur
 les années 1970. Nous découvrons que ces premiers travaux de modélisation
 des objets archéologiques à l’aide de « langages d’analyse » et de
 « domaines » n’ont rien à envier aux projets actuels à base de RDF et
 d’OWL, qu’ils recèlent même souvent une finesse d’analyse digne des
 projets les plus aboutis. Mais ces travaux sont surtout intéressants
 par les débats qu’ils firent naître dans la communauté des archéologues,
 débats d’une profondeur théorique suffisante pour rester, nous semble-t-il,
 applicables plus de 45 ans après. Parmi les critiques de l’époque, la
 plus intéressante et la plus constructive est probablement celle de
l’archéologue Philippe Bruneau  : contrairement aux objets de la Nature,
 les objets des Sciences humaines, parce qu’ils sont déjà porteurs de
sens, doivent être décrits avec des méthodes sémiotiques (plutôt que sémantiques).
 La prise en compte du contexte ne se traduit alors pas par un vague
principe ou par un modèle de plus, mais par un refus systématique des
modèles à portée générale. Les modèles doivent tenir compte du fait que
seuls un petit nombre des traits pouvant caractériser un objet seront
pertinents et ce par différence avec les objets qui sont proches de lui
 dans son univers technique. Consciemment ou inconsciemment, un certain
nombre de travaux d’aujourd’hui en ingénierie des connaissances
s’inscrivent au moins partiellement dans ces perspectives. Comme nous
l’illustrons avec nos propres logiciels et expérimentations, la prise
en compte de cette approche sémiotique trace des perspectives
prometteuses pour l’instrumentation de la pratique quotidienne des
chercheurs en Sciences humaines ainsi que pour la médiation scientifique.
 Mais, par un juste retour aux sources de l’intelligence artificielle
 (à l’époque où sa visée était moins substitutive que compréhensive)
 l’intérêt de la prise en compte de l’approche sémiotique est peut-être
 plus grand encore dans la foule de questions de conception qu’elle suscite,
 questions anecdotiques à première vue, mais liées finalement à ce qu’est
 le sens et à ses modes de construction.`},
	{"Conscience sans Cortex",
		"Michel Dojat,Manik Bhattacharjee,Christian Graff",
		"",
		"",
		`Avec l'avancée des connaissances en neurosciences, en éthologie et
le développement de l'intelligence artificielle, le conscience est-elle
toujours une spécificité humaine comme elle l'a été considérée par dans
les siècles passés ? A quel type de conscience les entités sans cortex
pourraient-elles prétendre ?`},

	{"Des intelligences Très artificielles",
		"Jean-Louis Dessalles",
		"",
		"",
		`Si vous marchez à reculons, les traces de pas que vous voyez devant
vous sont les vôtres. Aucun robot, aucune intelligence artificielle
(IA) ne sait ce genre de choses, sauf si l'on a pensé à les lui dire.
Les IA sont-elles si intelligentes que cela ? À bien y regarder, elles
apparaissent très intelligentes et très stupides à la fois. Pour quelle
raison ? En sera-t-il toujours ainsi ? Dans ce livre, Jean-Louis Dessall
s aborde ces questions d'une manière précise et accessible à tous.
Chaque lecteur trouvera dans ce livre de quoi le surprendre. Il nous
parle du passé, du présent et du futur des IA. Il évoque même ce qui,
selon lui, leur manque pour devenir... intelligentes.`},

	{"Un calcul cortical pour les systèmes situés",
		"Hervé Frezza-Buet",
		"",
		"",
		`Depuis les années 50 sont apparues, comme domaine d'étude, ce que l'on
appelle les sciences cognitives, qui ont fédéré des disciplines telles que
, entre autres, la neurobiologie, la psychologie, la philosophie,
la linguistique... et bien entendu l'informatique. Sans revenir sur
l'historique de la constitution des sciences cognitive, nous en
retiendrons qu'elles sont apparues à partir du moment où les sciences
sont devenues capables d'aborder la question de la conscience, en
décortiquant et objectivant les phénomènes de mémoire, de perception,
de langage et d'émotions. Objectiver le sujet, qui est au cœur des
sciences cognitives, est l'expression d'un paradoxe dont nous
parlerons dans ce chapitre, et la science informatique y a pris toute sa part.
L'Intelligence Artificielle (IA) a été le versant en informatique des sciences
cognitives. Parler d'intelligence artificielle fait d'ailleurs toujours
l'objet de polémiques, la question de savoir jusqu'où l'on peut dire qu'une
machine est capable d'intelligence n'étant pas tranchée aujourd'hui. Face à
cette indétermination, nous soutiendrons l'hypothèse dite de l'IA forte,
qui propose de considérer que l'ensemble de ce que l'on peut observer chez
l'homme en termes de langage, pensée, conscience, est strictement le fruit
de son métabolisme, neuronal en particulier. Nous ne demanderons pas au
lecteur d'adhérer à cette hypothèse, mais soulignons ici qu'elle motive
les travaux et les orientations de recherches que nous présentons dans ce
mémoire. L'hypothèse d'IA forte trouve un écho particulier en informatique
pour les raisons suivantes. L'informatique est une discipline qui dès
l'origine [Turing, 1936; Church, 1936] a su abstraire la mécanique du
traitement de l'information de son support. En effet, les machines de
Turing ont existé bien avant d'être instanciées sur support physique. Selon
cette perspective, qu'un traitement soit effectué sur silicium, au sein
d'une clepsydre améliorée, ou sur un support neuronal ne change rien à
l'affaire. La métaphore de la chambre chinoise de John Searle [Searle, 1980]
illustre cette indépendance au support. Elle compare l'ordinateur à une
personne enfermée dans une salle qui manipule des symboles chinois auxquels
elle ne comprend rien, suivant pour ce faire un système de règles. Si l'on
adhère à l'hypothèse d'IA forte, ce que ne fait pas Searle, le système de
règles pourrait décrire une intelligence, équivalente à celle de l'Homme,
dont l'opérateur est le moteur. Ce qui motive notre recherche informatique
est l'hypothèse de l'existence de ces règles, autrement dit, d'un programme
qui conduise 'a ce qu'une intelligence de même nature que celle de l'Homme
puisse être instanciée par une machine de Turing. Faire cette hypothèse,
toutefois, ne permet pas de guider la conception du programme, ou plutôt
d'un programme, qui puisse doter une machine d'intelligence. Il est alors
nécessaire de trouver ailleurs les arguments permettant de concevoir ces
programmes. Là encore, l'hypothèse d'IA forte joue un rôle central. Si on
admet que l'intelligence dont l'Homme fait preuve n'est que le résultat de
la mécanique de ses neurones, il devient pertinent de s'inspirer des neurones.`},

	{"Hanter la machine : reconquêtes de la conscience humaine",
		"Simon Bréan",
		"",
		"",
		`Le mouvement cyberpunk a fourni de multiples angles d’approche pour envisager
les rapports entre l’individu et des technologies informatiques modifiant en
profondeur son corps et ses modes d’être en société. Il n’a pas existé de
courant cyberpunk français à proprement parler, même si plusieurs œuvres
françaises notables en ont repris l’imagerie et les concepts dans un mouvemen
d’hommage ou d’appropriation. L’article vise à identifier une singularité
persistante dans le traitement du rapport humain-machine informatique au sein
du cyberpunk « à la française » : l’importance de la conscience humaine pour
la conception de l’intelligence artificielle.`},

	{"Le droit sous le règne de l’Intelligence Artificielle.",
		"Hervé Causse",
		"",
		"",
		`L'étude d'une trentaine de réalités sociales et juridiques fait penser que
les systèmes d'intelligence artificielle modifieront la façon de voir le droit
et d'en faire. Les problématiques utiles à l'analyse dépassent le domaine
juridique. Le droit de l'intelligence artificielle, à venir, exige en
préalable de mûrir des réalités sur lesquelles le juriste passe vite (le réel,
les langages, le code, l'objet, les systèmes, la conscience, la confiance
et bien entendu l'intelligence). Au moyen d'une méthode linéaire qui se
revendique de la liberté et du ""linéarisme"", escomptant sur la sérendipité,
l'examen de ces thèmes laisse apparaître une désarmante connexion entre
langage, système, science et intelligence. Elle signe la faiblesse des
langues actuelles et éclaire l'aptitude des systèmes d'IA, très rigoureux
à mieux les saisir pour tendre vers des analyses intelligentes. Le juriste
inquiet de ses métiers devrait s'inquiéter de sa méthode langagière qui
néglige les concepts - ou généraux ou de de la règle de droit.
Car les systèmes, eux, finiront par les appréhender comme ils ne le furent
jamais. Le changement dans le traitement de la langue annonce un choc
épistémologique dans toute les science. L'IA, qui n'avait pas besoin de
cela, se confirme en véritable mythe qui, comme telle, surplombe notre
civilisation. Dans sa quête de renouvellement, il se pourrait que le
juriste doive pousser jusqu'à la poésie, comme on y recourt parfois en
sciences exactes, pour dominer les énoncés des systèmes d'IA à partir
desquels, demain, on fera du droit.`},
	{"Intelligences artificielles, consciences réelles",
		"Gérard Dastugue",
		"",
		"",
		`Particularité du cinéma, l’écran « est un miroir où le spectateur
peut trouver d'autres identifications que son propre corps » (Christian
Metz). Il est à la fois un objet et un sujet, pouvant recevoir toutes
les projections sauf celle du corps même du spectateur. De cette absence
physique à l’écran (celui-ci étant un miroir symbolique et non tangible)
naît une identification non pas comme objet (la représentation de soi)
mais comme un avatar psychique (identification au personnage). Le lien
entre le spectateur et le sujet filmique serait-il du même genre que le
lien humain-cyborg ? Si l'homme augmenté s'écrit en termes hyperhumanistes,
il semblerait que la fiction cinématographique prenne le chemin de
l'inquiétante étrangeté et de la suspicion. La science-fiction au
cinéma semble donc préférer présenter l'intelligence artificielle
comme un danger et non une solution. Ces choix artistiques interrogent
finalement le cinéma sur sa propre nature, une malice de l'autre côté
du miroir.`},
	{"Nouvelles approches en Robotique Cognitive",
		"Mehdi Khamassi, Stéphane Doncieux",
		"",
		"",
		`Ce volume présente un ensemble de contributions montrant les évolutions
récentes des recherches du domaine dit de « Robotique Cognitive ».
Cette dénomination vaut dès lors que l’on cherche à faire réaliser
au robot des tâches qui semblent nécessiter chez l’homme la mise en
œuvre de fonctions cognitives telles que (sans être exhaustifs)
l’apprentissage, l’interaction sociale, la perception, la motricité,
la cognition spatiale, le raisonnement, le langage ou encore la
conscience. Le but du volume est de montrer en quoi les objectifs
et méthodes particuliers de ce domaine se sont historiquement
distingués d’autres travaux en Robotique ou en Intelligence Artificielle,
pour se rapprocher des autres disciplines des Sciences Cognitives.
En corollaire, le volume vise à expliciter certains des ponts
possibles qui peuvent permettre à ces disciplines de se féconder
mutuellement. En particulier, un des objectifs de ce volume est
d’illustrer en quoi l’expérimentation robotique peut servir de
plateforme de test d’hypothèses d’autres disciplines des Sciences
Cognitives, et ainsi contribuer à l’étude de la cognition biologique.`},
	{"Vers une robotique du traduire",
		"Anne Baillot, Ellen Carter, Thierry Grass, Pablo Ruiz",
		"",
		"",
		`L’apparition sur la toile en 2017 de nouveaux services de traduction
automatique neuronale reposant sur des algorithmes d’intelligence
artificielle comme DeepL et Google Translate correspond à un nouveau
bond en avant en matière de traduction automatique. Ces systèmes récents,
comme les systèmes de la génération précédente de traduction automatique
statistique et de traduction automatique statistique factorisée,
fonctionnent à partir de grands corpus alignés et produisent des
résultats dont la qualité est pour certains comparable à certaine
traductions humaines. Il s’ensuit que pour produire une valeur ajoutée,
le traducteur doit apporter un plus par rapport à la machine. Ce plus
peut être inhérent à certains domaines où l’emploi de la machine n’a
en soi guère d’intérêt du fait de la dimension essentiellement esthétique
de la traduction : c’est le cas des traductions littéraires qui, si
de nombreux travaux de traductologie universitaire y prennent appui,
ne représentent qu’une petite partie de l’activité de traduction
professionnelle existante. Comme de surcroît la machine permet des
gains de productivité de l’ordre de 150 à 200% (certains traducteur
atteignent des rendements de 6000 à 8000 mots par jour), la technique
de la post-édition tend à s’imposer de plus en plus dans les industrie
de la langue.`},
	{"Contribution des Sciences Sociales dans le domaine de l'Intelligence Artificielle Distribuée : ALONE, un modèle hybride d'agent apprenant.",
		"Isabelle Jars",
		"",
		"",
		`L'apprentissage fait partie des expériences qui stimulent et
structurent le savoir-faire individuel de tout être humain depui
sa naissance. Sa complexité et ses mystères suscitent l'intérêt de
nombreuses recherches dans des disciplines aussi variées que les
sciences sociales et l'informatique. En prenant conscience de cette
diversité d'approche et des liens pluridisciplinaires qui en découlent,
nous avons décidé de nous intéresser à cette problématique. L'objectif
visé est de modéliser le comportement d'apprentissage par l'intermédiair
des systèmes multi-agents en intégrant des recherches issues des
sciences sociales dans notre modèle. De ce fait, nous proposons
une modélisation (baptisée ALONE) qui allie les spécificités
théoriques des agents à la richesse des travaux socio-constructivistes
sur l'importance du langage et des interactions lors du processus
d'apprentissage humain. Notre contribution dans ce domaine est à
la fois méthodologique et applicative.`},
	{"L'homme, l'animal et la machine - Perpétuelles redéfinitions",
		"Georges Chapouthier, Frédéric Kaplan",
		"",
		"",
		`Les animaux ont-ils une conscience ? Les machines peuvent-elles se
montrer intelligentes ? Chaque nouvelle découverte des biologistes,
chaque progrès technologique nous invite à reconsidérer le propre de
l'homme. Ce livre, fruit de la collaboration entre Georges Chapouthier,
biologiste et philosophe de la biologie, et Frédéric Kaplan,
ingénieur spécialiste de l'intelligence artificielle et des interfaces
homme-machine, fait le point sur les multiples manières dont les animaux
et les machines peuvent être comparés aux êtres humains. Après un
panorama synthétique des capacités des animaux et des machines à apprendre,
développer une conscience, ressentir douleur ou émotion, construir
une culture ou une morale, les auteurs détaillent ce qui nous lie à
nos alter-egos biologiques ou artificiels : attachement, sexualité,
droit, hybridation. Au-delà, ils explorent des traits qui semblent
spécifiquement humains - l'imaginaire, l'âme ou le sens du temps -
mais pour combien de temps encore... Une exploration stimulante au
coeur des mystères de la nature humaine, qui propose une redéfinition
de l'homme dans son rapport à l'animal et à la machine`},
	{"Transition between cooperative and collaborative interaction modes for human-AI teaming",
		"Adrien Metge, Nicolas Maille, Benoît Le Blanc",
		"",
		"",
		`Avec l'introduction de l IA dans le pilotage des véhicules terrestres
ou aériens, la répartition des rôles entre opérateur et système devra
évoluer de manière dynamique. A travers une expérimentation en micro-monde
sur la supervision d'un drone intelligent, nous étudions comment
une telle transition entre modalités d'interaction coopératives et
collaboratives peut affecter l'expérience et le choix de l'opérateur.
Nous observons des variables comme le sentiment de responsabilité ou
la confiance et constatons que les opérateurs ont faiblement conscience
de l'influence de l'IA sr leur propre prise de décision.`},
	{"Un système Ingénieux de perception bio-inspiré basé sur les capacités auditives cognitives humaines",
		"Yu Su",
		"",
		"",
		`Développer une machine capable d’une perception consciente de
l’environnement dans lequel elle évolue, aux côtés et avec des
humains, est l’un des objectifs de l'intelligence artificielle
bio-inspirée (IAB). Les communautés des chercheurs en IA et en
IAB admettent généralement que l’adjonction d’une capacité artificielle
faisant émerger une sorte de « prise de conscience » ou un traitement
« conscient » de l’information par une machine conduirait vers une
technologie beaucoup plus puissante et plus avancée que celles basées
sur l’AI conventionnelle.L'ouïe est l’un des principaux systèmes
sensoriels du système cognitif humain. Les oreilles transforment
la myriade de stimulus perçus de l’environnement ambiant en signaux
(impulsions) nerveuses générées par différents types de cellules
nerveuses et cela à tout moment, même lorsque nous nous endormons.
En effet, avec et aux côtés de la vision (i.e. capacité visuelles),
le système auditif constitue un sens fondamental de la perception chez
l’humain. Motivé par l’importance du complément auditif chez l’humain
dans la perception et la caractérisation par ce dernier de l’environnement
dans lequel il évolue et compte tenu des limites actuelles pour la
simulation du mécanisme cognitif auditive humain, l’objectif principal
de ce travail doctoral est de fournir aux machines une capacité auditive
cognitive artificielle dotant ces dernières d’une perception augmentée et
adaptée de l'environnement à l’image de celle développée chez les humains.
Pour atteindre cet objectif, tout d’abord, une étude des travaux de
recherche les plus récents, couvrant les modèles d’attention auditive,
les techniques de classification du son environnemental, celles basées
sur l’apprentissage profond (deep-learning) et les mécanismes de réponse
auditive humaine, a été effectuée permettant de mieux comprendre l’état
actuel de l’art et la complexité de la réalisation des objectifs visés
par le présent travail doctoral. Cette étude a mis en exergue les
insuffisances inhérentes aux techniques existantes et a orienté nos
investigations vers une modélisation des mécanismes bio-inspirés de la
détection de la divergence auditive. Ces modèles ont été associés aux
réseaux de neurones convolutionnels (CNN) pour catégoriser les sons détecté
dans l’environnement en exploitant un système à base de connaissances.
Ensuite, les travaux ont conduit à la mise en œuvre d’un modèle pour
la détection de la déviance auditive en utilisant à la fois des caractéristiques
temporelles et spatiales du son perçu (domaines temporel et spatial).
Une approche d’extraction de ce type de caractéristiques a été proposée.
Ainsi, les caractéristiques précitées contribuent à la détection de la
déviance et de la saillance auditive dans chaque domaine (i.e. domaine
temporel et domaine spatial) pour, ensuite être combinées afin de
fiabiliser la détection et la catégorisation du son perçu de l'environnement réel
(i.e. le résultat final). Les résultats expérimentaux montrent la viabilité
du modèle proposé pour détecter des sons saillants déviants dans un clip
sonore ainsi que la robustesse et une précision des modèles proposés.
Finalement, les travaux ont conduit à la mise au point d’un modèle
puissant de détection et caractérisation des sons environnementaux,
issu d’une fusion de deux CNN à 4 couches.`},

	{"Construction et conceptualisation de connaissances en robotique autonome",
		"Cristiano Russo",
		"",
		"",
		`L'emploi de robots personnels, ou de robots de service, a suscité
beaucoup d'intérêt ces dernières années avec une croissance étonnante
de la robotique dans différents domaines. Concevoir des robots compagnons
capables d'assister, de partager et d'accompagner des personnes à autonomie
limitée dans leur vie quotidienne est le défi de la décennie à venir.
Cependant, les performances des systèmes robotisés et des prototypes
actuels sont très loin de répondre à un tel défi. Bien que des robot
humanoïdes sophistiqués aient été développés, de nombreux efforts sont
nécessaires pour améliorer leurs capacités cognitives.En effet,
les robots (ainsi que les prototypes) disponibles dans le commerce
ne sont pas encore capables de s'adapter naturellement à l'environnement
complexe dans lequel ils sont censés évoluer avec les humains.
De la même façon, les prototypes existants ne sont pas en mesure
d'interagir de plusieurs manières avec leurs utilisateurs. En fait
ils sont encore très loin d'interpréter la diversité et la complexité
des informations perçues ou encore de construire des connaissances
relatives au milieu environnant. Le développement d'approches bio-inspirées
basées sur la cognition artificielle pour la perception et l'acquisition
autonome de connaissances en robotique est une méthodologie appropriée
pour surmonter ces limites. Un certain nombre d'avancées ont déjà permis
de réaliser un système basé sur la cognition artificielle permettant
à un robot d'apprendre et de créer des connaissances à partir de ses
observations (association d'informations sensorielles et de sémantique
naturelle). Dans ce contexte, le présent travail tire parti du
processus évolutif d'interprétation sémantique des informations sensorielles
pour faire émerger la conscience de la machine sur son environnement.
L'objectif principal de la thèse de doctorat est de poursuivre les efforts
déjà accomplis (recherches) afin de permettre à un robot d'extraire,
de construire et de conceptualiser les connaissances sur son environnement.
En effet, la motivation de cette recherche doctorale est de généraliser
les concepts précités afin de permettre une construction autonome,
ou semi-autonome, de connaissances à partir de l'information perçue
(par exemple par un robot). En d'autres termes, l'objectif attendu
de la recherche doctorale proposée est de permettre à un robot de
conceptualiser progressivement l'environnement dans lequel il évolue
et de partager les connaissances construites avec son utilisateur.
Pour cela, une base de connaissances sémantique-multimédia a été
créée sur la base d'un modèle ontologique et implémentée via une bas
de données de graphes NoSQL. Cette base de connaissances est la pierr
angulaire du travail de thèse sur lequel de multiples approches ont
été explorées, basées sur des informations sémantiques, multimédia
et visuelles. Les approches développées combinent ces informations à
travers des techniques classiques d'apprentissage automatique, à la
fois supervisées et non supervisées, ainsi que des techniques
d'apprentissage par transfert pour la réutilisation de caractéristiques
sémantiques à partir de modèles de réseaux de neurones profonds.
D'autres techniques basées sur les ontologies et le Web sémantique
ont été explorées pour l'acquisition et l'intégration de nouveaux
savoirs dans la base de connaissances développée. L'étude de ces
différents domaines à conduit à la definition d'un modèle compréhensif
de gestion de la connaissance intégrant des caractéristiques relatives
à la perception et à la sémantique, qui peut également être utilisée
sur des plateforme robotiques. Les expériences menées ont montré une
correspondance efficace entre les interprétations basées sur des
caractéristiques sémantiques et visuelles, d'où la possibilité pour
un agent robotique d'élargir ses compétences de généralisation des
connaissances dans des environnements encore inconnus (voire partiellemen
connus), ce qui a permis d'atteindre les objectifs fixés.`},
	{"Comportements et Mémoires",
		"Nicolas P. Rougier",
		"",
		"",
		`La survie, pour la majeure partie du règne animal, dépend directement
de la capacité à se mouvoir dans un environnement (connu ou inconnu)
afin d'être en mesure de rallier, de façon précise et sûre, des lieux
spécifiques tels que l'habitat, un point d'eau ou bien encore un
lieu de nourriture. Les techniques mises en oeuvre pour satisfaire
ces buts sont généralement de nature très diverse, allant des plus
simples aux plus élaborées selon les espèces. Ainsi, certaines
espèces utilisent l'orientation par rapport au soleil, d'autres
utilisent les phéromones et d'autres encore utilisent un plan du
métro. Toutefois, si les différentes études menées sur la navigation
animale ont permis de mettre en évidence un large éventail de stratégies
autorisant la navigation autonome, la nature très diverse des mécanismes
impliqués peut être appréhendée selon une taxonomie à 4 niveaux proposé
par Trullier et col. que nous détaillerons. Or, si les techniques les
plus simples peuvent être utilisées relativement simplement au sein d
modèles de navigation, les techniques plus élaborées telles que
celles impliquant la notion de carte cognitive requièrent des modèles
plus complexes manipulant à la fois des notions déclaratives telles
que les lieux et des notions procédurales telles que la coordination
sensori-motrice permettant par exemple le passage d'un lieu à un autre.
Nous proposons dans cet exposé de caractériser les systèmes mnésiques
impliqués sur la base d'observations psychologiques et neurologiques.
En effet, les concepts de mémoire déclarative et procédurale ont été
proposés par Cohen et Squire en 1980 (2) sur la base d'un vocabulaire
déjà présent en Intelligence Artificielle. l'information stockée en
mémoire déclarative serait accessible à la conscience et pourrait être
utilisée via le langage ou des images mentales. Les informations sont
des connaissance de type général tels que des faits, des évenements,
etc (par exemple retenir un numéro de téléphone). La mémoire procédurale
ne serait elle pas directement accessible à la conscience et permettrait
d'acquérir des aptitudes indissociables de l'action (par exemple savoir
faire du vélo) et ne pourrait donc s'exprimer qu'au cours d'une action
Nous souhaitons notamment souligner comment le connexionnsime peut
promouvoir une coopération entre ces deux systèmes de mémoires en évitant
l'écueil de l'approche ""boîte noire"" et nous présenterons à cet
effet quelques modèles de mémoire procédurale et déclarative.`},
	{"La cognition : du neurone a la societé",
		"Daniel Andler, Thérèse Collins, Catherine Tallon-Baudry",
		"",
		"",
		`La cognition désigne l'ensemble des phénomènes qui se rapportent à
l'esprit humain, son fonctionnement, ses effets sur le comportement,
son émergence au cours de l'évolution, et son développement, typique
ou non. Elle inclut les grandes facultés identifiées par la tradition
(attention, mémoire, raisonnement, décision, langage...), mais
tout autant la perception, la motricité, les émotions, la
conscience, la socialité. Rendre compte de cet immense domaine es
le but des sciences cognitives, fédération de disciplines allant d
la biologie à la linguistique, de la psychologie à l'intelligence
artificielle, de la philosophie à l'anthropologie.`},

	{"Délibérer avec l'intelligence artificielle au service de l'intelligence naturelle.",
		"Frédéric Alexandre, Thierry Viéville, Marie-Hélène Comte",
		"",
		"",
		`La numérisation de la société et le traitement automatique de
l’information, y compris avec des techniques dites d’intelligence
artificielle, induisent des ruptures dans notre façon de penser,
calculer et délibérer. Mais comment fonctionnent ces fonctions cognitives
et comment peuvent-elles être affectées par cette révolution numérique ?
Pour comprendre en profondeur cela, nous allons d’abord prendre
le temps d’expliquer ce que nous comprenons aujourd’hui de notre
intelligence biologique qui pense, ce qui offrira un éclairage
crucial sur ce qui se passe quand on utilise des machines qui calculent,
avant de conclure en quoi cela aide à réfléchir sur comment délibérer
Car oser comprendre les aspects scientifiques et techniques de la pensée
et du calcul est essentiel pour se donner les moyens de délibérer en
toute conscience avec les outils intellectuels et numériques qui nous
sont donnés.`},

	{"Psychologie des êtres artificiels",
		"Pierre Crescenzo",
		"",
		"",
		`L'émergence de l'Intelligence Artificielle comme thème de recherche
il y a quelques dizaines d'années a déjà mené à de nombreuses réflexions
complexes et parfois dérangeantes. Plus récemment, des réalisations
concrètes que l'on peut dire intelligentes s'insèrent de plus en plus
au sein des sociétés humaines en les remettant en question, parfois
jusque dans leurs principes essentiels et leurs bases. Ces IA, robots,
objets connectés sont désormais des éléments constitutifs de nos vies.
Et nous les percevons de plus en plus souvent comme des êtres réellement
intelligents et autonomes. Nous interagissons avec eux, avec la
perception de leur intelligence, de leur autonomie et de leur psychologie.
Pour combien de temps encore ne s’agira-t-il que d’une perception, et
non du début d’une nouvelle réalité ? Nul ne le sait. Mais il est important
de se poser la question de leur existence en tant qu’êtres, qu’êtres pensants
d'avoir conscience que ces questions sont importantes, pour pouvoir
être des informaticiens responsables, des spécialistes d'intelligence
artificielle éclairés et ouverts aux futurs possibles. Cette ouverture
est l'objectif immense de ce modeste article.`},

	{"L'intelligence artificielle : entre opportunités et risques légitimes",
		"Marie Noeline Sinapin",
		"",
		"",
		`Cet article s’intéresse à l'intelligence artificielle, perçu comme un
progrès majeur depuis le début 2016 et permet de concevoir des systèmes
de plus en plus sophistiqués, au point de rêver aux voitures sans
conducteur, une vision imaginée par Volvo et d'autres constructeurs
automobiles. Les innovations sont nombreuses grâce aux récents progrès
de l'intelligence artificielle, et engendreraient des répercussions insoupçonnées
dans tous les domaines de nos activités terrestres. Si les entreprises ont
pris conscience de l'intérêt de l'I.A, très peu connaissent ses techniques
d'apprentissage, ses répercussions et ses risques.`},

	{"L’avocat et l’algorithme : quelles transformations des compétences pour la profession ?",
		"Philip Milburn",
		"",
		"",
		`L'article s'interroge dans un premier temps sur le travail et les compétences
réels des avocats, au-delà des discours et de la rhétorique de la profession.
Cet examen en détail de la nature de cette activité concrète permet de pose
quelques linéaments quant à la manière dont les services juridiques offert
par l'intelligence artificielle sont susceptibles de venir l'impacter, d'après
les publications de Richard Susskind. Afin d'éviter une disqualification
économique de leur expertise, les avocats doivent prendre conscience et
mettre en valeur la part non juridique de leur compétence.`},
	{"Reconnaissance des formes et vision par ordinateur",
		"Peter Sturm, Serge Garlatti",
		"",
		"",
		`Face à l'essor grandissant des sciences du traitement de l'information e
la prise de conscience de l'importance de l'interdisciplinarité, il est plus
que jamais nécessaire de mettre en avant les spécificités et les
complémentarités des deux domaines RF et IA, reconnaissance des formes et
intelligence artificielle. Dans un contexte où les défis rencontrés impliquent
notre capacité à créer des synergies entre les domaines et les disciplines,
RFIA est précisément l'occasion d'affronter les enjeux pour relever ces défis.
Depuis 1977, la conférence RFIA est le rendez-vous incontournable des scientifiques,
chercheurs et professionnels qui animent les domaines de la reconnaissance des
formes et de l'intelligence artificielle. Depuis sa naissance, la conférence a
su valoriser les recherches des chercheurs de renommée internationale mais aussi
repérer les travaux prometteurs des jeunes chercheurs, offrant un panorama très
représentatif de l'état de l'art comme des perspectives les plus stimulantes.`},

	{"Le Numérique et la vie",
		"Arthur Guezengar",
		"",
		"",
		`Que peut-on attendre de la philosophie face aux réalités technologiques ?
Si elle n'a jamais été insensible à l'agir technique, plus que jamais la
pensée philosophique se trouve requise du fait de la puissance de transformation
du monde manifestée aujourd'hui par la technologie. Ce volume se donne pour
premier objectif minimal de reconstituer la faculté de juger mise à mal ou
brouillée par les évolutions technologiques. Il est également possible que
les nouveaux champs technologiques émergents, de par l'originalité et la
capacité de reconfiguration des inventions qu'ils connaissent, conduisent
bientôt la philosophie à proposer de nouvelles sous-disciplines, en l'obligeant
à travailler sur elle-même. Semblent, en ce moment même, susceptible de
provoquer ce genre d'évolutions, la combinaison de l'informatique et du
numérique ou l'Intelligence Artificielle, ses déclinaisons potentiellement variées
à l'infini mais déjà sensibles dans la robotique, la transformation de la
production par exemple sous l'effet de l'impression 3D, l'interaction avec
des êtres artificiels ou synthétiques de tous ordres et de toutes tailles,
qu'il s'agisse des agents conversationnels, des drones ou des smart cities.
De telles évolutions en cours font qu'un des attendus de la démarche philosophique
peut concerner une vaste ambition, à savoir celle de dessiner les cadres
d'un environnement désirable non seulement pour les humains, mais pour
toutes les êtres vivants et existants, qu'ils soient naturels ou artificiels.
En ce cas, la philosophie appliquée à la technique d'aujourd'hui se donne
une ambition normative de très haut niveau en accompagnant le déploiement
des nouvelles formes de conscience sensibles à la nature et aux vivants ou
agissants non-humains. Les contributions comprises dans ce dossier portent
sur l'éthique et la philosophie politique appliquée aux activités humaine
telles qu'elles sont aujourd'hui assistées et transformées par la technique
faisant système et valant déjà comme vision du monde. En confrontant la
philosophie au terrain des innovations et en observant les pratiques, les
autrices et les auteurs de ce dossier s'emploient à déterminer les modalités
de l'éthique et de la politique appliquée à l'activité humaine telle qu'elle
est indissociablement assistée et régie par la technologie.`},

	{"L’intelligence naturelle : la science des données de l’IA",
		"Arnaud Martin",
		"",
		"",
		`L'intelligence artificielle est devenue aujourd'hui un enjeu économique
et se retrouve au cœur de nombreuses actualités. En effet les applications
d'un grand nombre d'algorithmes regroupés sous le terme d'intelligence
artificielle se retrouvent dans plusieurs domaines. Citons par exemple
l'utilisation de l'intelligence artificielle dans les processus de
recrutement ou encore pour le développement de voitures autonomes. Dans
ce contexte, la France veut jouer un rôle dans lequel le rapport Villani
pose les bases d'une réflexion. Ces méthodes et algorithmes d'intelligence
artificielle soulèvent cependant des craintes et des critiques liées à
la perte du sens commun issu des relations humaines. Il faut toutefois
avoir conscience que l'ensemble de ces approches ne font qu'intégre
des connaissances humaines qui ont été au préalable modélisées bien
souvent à partir de processus d'extraction de connaissances. C'est
dans ce cadre que nous menons des recherches dans l'équipe DRUID
(Declarative & Reliable management of Uncertain, user-generated
Interlinked Data) de l'IRISA (Institut de recherche en Informatique
et Systèmes Aléatoires).`},

	{"Peut-on modéliser la conscience à l'aide d'un système informatique ?",
		"Bernard Victorri",
		"",
		"",
		`Pour répondre à cette question, on s'attache d'abord à préciser non
seulement ce que l'on entend par conscience, mais aussi le sens précis
que l'on donne à modéliser. Certains modèles sont conçus avec une finalité
propre, qui ne doit rien au système modélisé, sinon une inspiration
qui a conduit à en utiliser certaines caractéristiques. A l'autre extrême,
d'autres modèles sont de véritables démonstrations constructives : en
exhibant un système qui fonctionne comme le système étudié, en respectant
aussi bien ce que l'on connaît de son organisation interne que ce que
l'on peut observer de ses comportements, on prouve la cohérence et la
complétude de la théorie qui a permis de construire ce modèle. La grande
majorité des modèles de la conscience se situe entre ces deux extrêmes.
Par ailleurs, indépendamment de l'existence de ces modèles plus ou moins
explicatifs de la conscience humaine, le problème peut se poser de savoir
si une machine « intelligente » est susceptible d'être le siège de véritables
phénomènes conscients : bien que cette question fasse l'objet de vifs débats,
on peut douter de leur intérêt en l'état actuel de nos connaissances.`},
	{"Intelligence artificielle: l'apport des paradigmes incarnés",
		"Pierre de Loor, Alain Mille, Mehdi Khamassi",
		"",
		"",
		`Cet article a un double objectif : Le premier est de présenter les
différentes propositions relatives à l'approche incarnée de la cognition
faites par la communauté informatique et robotique. Le deuxième es
de mener un débat sur leurs apports et leurs limites, relativement
aux questions les plus délicates des sciences cognitives que sont la
construction du sens, la conscience phénoménale ou encore les liens
entre esprit, matière et organisation. La première partie de l'article
dresse un rappel historique des objectifs initiaux de l'intelligence
artificielle ainsi que les différentes orientations qui sont désormais
prises par cette communauté. La partie suivante positionne le débat à la
fois sur les questions fondamentales que peut étudier ou pas l'intelligence
artificielle pour répondre aux questions difficiles des sciences de la
cognition et en particulier les intérêts ou limites liés à l'utilisation
d'une approche incarnée pour y répondre. La troisième partie consiste à
détailler l'approche incarnée selon une structuration en familles, définies
par des domaines ou des focalisations différentes en neurosciences,
psychologie ou biologie. Nous faisons une description des principes
sur lesquelles chacune d'elles repose et nous en identifions les limites
et les possibilités relativement au débat posé. Le tout est synthétisé
par une conclusion mettant en perspectives les recherches présentées.`},
}

var rnd *rand.Rand

var articleAll []Post

func genDate(last time.Time, interval int) time.Time {
	nbDay := rnd.Intn(interval)
	res := last.AddDate(0, 0, -nbDay)
	return res
}
func genChoice(sList []string) string {
	return sList[rnd.Intn(len(sList))]
}

// Génère des post sans answers pour chaque article
func main() {

	var err error
	var date time.Time

	rnd = rand.New(rand.NewSource(time.Now().Unix()))

	// articles de Mandrake, commence par le dernier
	var lastDate time.Time
	for ida := len(artMathison) - 1; ida >= 0; ida = ida - 1 {
		art := artMathison[ida]

		var date time.Time
		if art.DatePost != "" {
			date, err = time.Parse("2006-01-02T15:04:05", art.DatePost)
			if err != nil {
				fmt.Printf("FATAL cannot parse time: %v\n", err)
				panic(err)
			}
			lastDate = date
		} else {
			date = genDate(lastDate, 30)
		}

		post := Post{
			Server:  addrServer,
			Group:   groupMathison,
			Date:    date,
			Author:  art.AuthorPost,
			Subject: art.Title,
			Content: fmt.Sprintf("Authors: %s\n\nRésumé: %s\n", art.Authors, art.Abstract),
		}
		articleAll = append(articleAll, post)
	}

	// then other articles
	lastDate, err = time.Parse("2006-01-02T15:04:05", lastDateOther)
	if err != nil {
		fmt.Printf("FATAL cannot parse time: %v\n", err)
		panic(err)
	}

	for _, art := range artDivers {
		if art.DatePost != "" {
			date, err = time.Parse("2006-01-02T15:04:05", art.DatePost)
			if err != nil {
				fmt.Printf("FATAL cannot parse time: %v\n", err)
				panic(err)
			}
		} else {
			date = genDate(lastDate, 700)
		}
		if art.AuthorPost == "" {
			art.AuthorPost = genChoice(logDivers)
		}
		post := Post{
			Server:  addrServer,
			Group:   groupMathison,
			Date:    date,
			Author:  art.AuthorPost,
			Subject: art.Title,
			Content: fmt.Sprintf("Authors: %s\n\nRésumé: %s\n", art.Authors, art.Abstract),
		}
		articleAll = append(articleAll, post)

	}

	yamlArt, err := yaml.Marshal(articleAll)
	if err != nil {
		panic(err)
	}
	fmt.Printf("## Articles *****\n%s\n", yamlArt)
}
