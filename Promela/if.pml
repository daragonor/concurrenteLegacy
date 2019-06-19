active proctype P() {
    byte a = 5
    byte b = 5
    
    if
    :: a >= b -> printf(":)\n")
    :: b >= a -> printf(":(\n")
    fi
    
    if
    :: a > 100 -> printf(":0\n")
    :: b > 100 -> printf(":S\n")
    :: else -> printf("NONONONONO\n")
    fi
    
    if
    :: a > 10 -> printf(":0\n")
    :: b > 10 -> printf(":S\n")
    fi
}