

export default function SlicesToSlices(data){
    if (typeof data !== Array || isNaN(data)||data.length<=0){
        return [];
    }
    slicesColl =new Array(Object.keys(data).length);
    for(let i =0;i<data.length;i++){
        let keys= Object.keys(data[i]);
        if (keys.length!=slicesColl.length){
            console.log("convert error",data);
            return [];
        }
        for (let j=0;j<keys.length;j++){
            if (slicesColl[j]==undefined||slicesColl[j]==null){
                slicesColl[j]=[];
            }
            slicesColl[j].push(data[i][j]);
        }
    }
    return slicesColl;
}