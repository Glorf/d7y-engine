# Engine
## Wczytywany plik mapy json (przykład)
```
{regions: [
    {
        name: "Brest",
        type: land,
        shortcut: "Bre", 
        adjacent: [
            {name: "Wales"},
            {name: "Paris"},
            {name: "Mid Atlantic"},
            {name: "English Channel"}
        ]
    },
    {
        name: "West Mediterraean",
        type: sea,
        adjacent: [
            {name: "Spain"},
            {name: "North Africa"},
            {name: "Gulf of Lyon"},
            {name: "Tunisia"},
            {name: "Tyrhennian Sea"}
        ]
    }
]}
```
## Automat stanów: (dla iteracji tygodniowej)
1) WIOSNA
    * Faza dyplomacji i pisania rozkazów (2 dni) - przyjmowanie rozkazów i validacja w locie
    * W razie braku dosłania rozkazu - wszystkie jednostki hold
    * Zamrożenie wysyłania na godzinę przed rozegraniem?
    * Rezolucja rozkazów - aktualizacja planszy
    * Faza odwrotu (1 dzień) - przyjmowanie rozkazów i validacja w locie
    * W razie braku dosłania rozkazu - pokonane armie są rozwiązywane. Aktualizacja planszy
1) JESIEŃ
    * Faza dyplomacji i pisania rozkazów (2 dni) - przyjmowanie rozkazów i validacja w locie
    * W razie braku dosłania rozkazu - wszystkie jednostki hold
    * Zamrożenie wysyłania na godzinę przed rozegraniem
    * Rezolucja rozkazów - aktualizacja planszy
    * Faza odwrotu (1 dzień) - przyjmowanie rozkazów i validacja w locie
    * W razie braku dosłania rozkazu - pokonane armie są rozwiązywane
    * Weryfikacja stanu gry, aktualizacja planszy i naliczanie punktów zwycięstwa
1) Rozpoczęcie kolejnej WIOSNY


## Parser i validator rozkazów:
* przyjmowanie ignore case, zarówno zapis skrócony jak i pełen
* usuwanie myślników i zastępowanie ich spacjami
1) Rozkaz Hold - domyślne działanie, nadpisuje poprzedni rozkaz dla jednostki\
np. `F London Holds` / `F Lon-Holds`
1) Rozkaz Move - próba ruchu armii\
np. `A Paris-Burgundy` / `A Par-Bur`
1) Rozkaz Support\
np. `A Gas S A Mar-Bur` / `A Gascony S A Marseilles-Burgundy`
1) Rozkaz Convoy \
np. `F GoL C A Spa-Nap` / `F Gulf of Lyon C A Spain-Naples`

## Frontend mailowy
1) Problem z ewaluacją gracza, widzę dwie opcje:
    * gracz podaje w mailu hasło które jednoznacznie go identyfikuje - bardziej 'wojennie'
    * zmusimy ich do korzystania z PGP i podpisywania maili kluczami - IMO not doable z nietechnicznymi
1) Decyzja co do zastosowanego interfejsu (moje dwa pomysły):
    * wszystkie rozkazy są wysyłane razem jednym mailem i natychmiast weryfikowane - dokładna wersja "komputerowa"
    rzeczywistej gry
    * wersja bardziej IMO realistyczna: każdy region ma własny adres email; wysłany rozkaz wirtualnie "trafia" do 
    konkretnego dowódcy, który sprawdza
    tylko jego poprawność i nie weryfikuje go względem rozkazów innych dowódców. W tym trybie gracze nie mogą widzieć
    swoich planowanych ruchów na mapie zanim nie minie tura. Jeśli gracz się pomyli i wyśle rozkazy do
    regionu przeciwnika, mail zostaje przechwycony i przekazany do odpowiedniego gracza. Mail nie dochodzi natychmiast:
    jego czas reakcji zależy od odległości regionu od stolicy oraz dostępności drogi lądowej lub morskiej. Potwierdzenie
    odebrania rozkazu też dociera dopiero po jakimś czasie. (chociaż ten feature miałby bardziej zastosowanie w grze realtime)


# Mapa graficzna
(janek oceń) -> pewnie zdefiniowane punkty na svg i kolorowanie ich?
