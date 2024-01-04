import { BorderBox1 } from '@jiaminghi/data-view-react';
import EChartsReact from 'echarts-for-react';
import React,{Component} from 'react'

const invalidResponse =[["ERC20","ERC721"],[{"name":"ERC20","value":100},{"name":"ERC721","value":50}]];

class TokenType extends Component{
    constructor(props){
        super(props);
        this.state={
            tokenEcharts:{
                color:{
                    type:'linear',
                    x:0,
                    y:0,
                    x2:1,
                    y2:0,
                    colorStops:[
                        {
                            offset:0,
                            color:'#d000d0'
                        },
                        {
                            offset:1,
                            color:'#7006d9'
                        }
                    ],
                    globalCoord: false,
                },
                    tooltip:{
                        trgger:'axis',
                        axisPointer:{
                            type:'shadow'
                        }
                    },
                    grid:{
                        left:'3%',
                        right:'4%',
                        bottom:'3%',
                        containLabel:true
                    },
                    xAxis:[
                        {
                            type:'value',
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
                        }
                    ],
                    yAxis:[
                        {
                            type:'category',
                            data:[],
                            axisTick:{
                                alignWithLabel:true
                            },
                            axisLabel:{
                                show:true,
                                textStyle:{
                                    color:'#c3dbff',
                                    fontSize:12
                                }
                            }
                        }
                    ],
                    series:[
                        {
                            name:"调用次数",
                            type:'bar',
                            barWidth:'60%',
                            data:[],
                        }
                    ]
                }
        };
    }
    asyncGetTokenTypes(){
        let result =undefined;
        try{
            throw("bad request")
        }catch(error){
            result=invalidResponse;
        }finally{
            return result;
        }
    }
    componentDidMount(){
       let [tokenType,tokens]= this.asyncGetTokenTypes();
       this.setState(preState=>({
        tokenEcharts:{
            ...preState.tokenEcharts,
            series:{
                ...preState.tokenEcharts.series,
                data:tokens,
            },
            yAxis:{
                ...preState.tokenEcharts.yAxis,
                data:tokenType
            }
        }
       }))
    }
    render(){
        return( 
            <div className='xpanel-wrapper xpanel-wrapper-6' style={{position:'relative'}}>
            <div className='content_title'>Token类型占比</div>
           <BorderBox1>
            <div className='xpanel'>
                <div className='fill-h'>
                    <EChartsReact  style={{width:'100%',height:'100%'}}  option={this.state.tokenEcharts}></EChartsReact>
                </div>
            </div>
           </BorderBox1>
          </div>
        )
    }
}


export default TokenType;