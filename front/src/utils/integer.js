

var zeroBytes = ['0','0','0','0','0','0','0','0','0','0','0','0']

const defaultMaxSliceLength = 12

export  function NumberToByteSlice(number){
    if (!isNumber(number)){
        return zeroBytes;
    }
    let result =new Array(defaultMaxSliceLength);
    var reverseIndex = defaultMaxSliceLength-1;
    while(number/10!=0 || reverseIndex>=0){
        if (number/10 == 0){
            result[reverseIndex] = 0;
            reverseIndex--;
            continue;
        }
        result[reverseIndex] = number%10;
        reverseIndex--;
        number= parseInt(number/10);
    }
    console.log(result);
    return result;
}

function isNumber(value) {
    return typeof value === 'number' && !isNaN(value);
  }