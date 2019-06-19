active proctype P() {

    int turn = 1
    
    if
    :: (turn == 1) ->
        turn = 2
        printf(turn)
    :: (turn == 2) ->
        turn = 1
        printf(turn)
    if

}