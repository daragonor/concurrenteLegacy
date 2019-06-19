byte n

proctype P() {
    byte i, temp;
    for (i : 1..10){
        temp = n;
        n = temp + 1;
    }
}


init {
    atomic{
        run P()
        run P()
    }
    
    (_nr_pr == 1) -> printf("N = %d\n", n);
    assert(n > 2)
}