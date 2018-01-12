from euler import timer, product


@timer
def euler8():
    number = '731671765313306249192251196744265747\
    4235534919493496983520312774506326239578318016\
    9848018694788518438586156078911294949545950173\
    7958331952853208805511125406987471585238630507\
    1569329096329522744304355766896648950445244523\
    1617318564030987111217223831136222989342338030\
    8135336276614282806444486645238749303589072962\
    9049156044077239071381051585930796086670172427\
    1218839987979087922749219016997208880937766572\
    7333001053367881220235421809751254540594752243\
    5258490771167055601360483958644670632441572215\
    5397536978179778461740649551492908625693219784\
    6862248283972241375657056057490261407972968652\
    4145351004748216637048440319989000889524345065\
    8541227588666881164271714799244429282308634656\
    7481391912316282458617866458359124566529476545\
    6828489128831426076900422421902267105562632111\
    1109370544217506941658960408071984038509624554\
    4436298123098787992724428490918884580156166097\
    9191338754992005240636899125607176060588611646\
    7109405077541002256983155200055935729725716362\
    69561882670428252483600823257530420752963450'
    print(max(product(number[i:i+13]) for i in range(1000-13)))

euler8()