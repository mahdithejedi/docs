Feature: stationExample

    Find areas with minimum amount of stations

Scenario Outline: SimpleSmallCityTest
    Given a <Dictionary> of stations with <InputList> of state needed
    When we insert them
    Then should return minium Amount of stations as <OutputList>
    Examples: station data
        | Dictionary | InputList | OutputList|
        | {'kone': {'id', 'nv', 'ut'}, 'ktwo': {'id', 'mt', 'wa'}, 'kthree': {'ca', 'nv', 'or'}, 'kfour': {'nv', 'ut'}, 'kfive': {'az', 'ca'}} | mt,wa,or,id,nv,ut,ca,az | ktwo,kthree,kone,kfive |

