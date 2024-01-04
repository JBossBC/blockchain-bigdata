import React,{Component} from 'react'

import { BorderBox8 } from '@jiaminghi/data-view-react';

class TransactionState extends Component{
  constructor(props){
    super(props);
    this.state={
      transactionStateEcharts:{
        title:{
          show:true,
          text:"近6个月交易趋势",
          x:'center',
          textStyle:{
            fontSize:14,
            fontStyle:'normal',
            fontWeight:'normal',
            color:'#01c4f7'
          }
        },
        tooltip:{
          trigger:'axis',
          axisPointer:{
            type:'shadow'
          }
        },
        legend:{
          data:["交易成功数(笔)","交易失败数(笔)"],
          textStyle:{
            fontSize:12,
            color:'#ffffff'
          },
          top:20,
          itemWidth:20,
          itemHeight:12,
          itemGap:10
        },
        grid:{
          left:'3%',
          right:'4%',
          bottom:'3%',
          containLabel:true
        },
        xAxis:{
          type:"category",
          data:[],
          splitLine:{
            show:true,
            lineStyle:{
              color:['#07234d']
            }
          },
          axisLabel:{
            show:true,
            textStyle:{
              color:'#c3dbff',
              fontSize:12
            }
          }
        },
        yAxis:{
          type:'value',
          boundaryGap:[0,0.01],
          splitLine:{
            show:true,
            lineStyle:{
              color:['#07234d']
            }
          },
          axisLabel:{
            show:true,
            textStyle:{
              color:'#c3dbff',
              fontSize:12
            }
          }
        },
        series:[{
          name:"交易成功数(笔)"
        }]
      }
    }
}

    render(){
        return(
        <div style={{width:"50%",paddingRight:8,position:'relative'}}>
        <BorderBox8 dur={8}>
          <div className='xpanel'>
            <div className='fill-h'></div>
          </div>
        </BorderBox8>
      </div>
        )
    }
}






export default TransactionState ;