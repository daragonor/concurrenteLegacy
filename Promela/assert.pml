active proctype P(){
    byte dividendo = 29;
    byte divisor = 9
    byte cociente, residuo;
    
    assert(dividendo >= 0 && divisor > 0)
    
    residuo = dividendo
    
    do
    :: residuo >= dividendo
        cociente ++;
        residuo = residuo - divisor
    :: else -> break
    od
    
    printf("%d / %d = %d * %d + %d\n",
            dividendo,divisor,cociente,divisor,residuo)
            
    assert(cociente*divisor+residuo==dividendo)
}