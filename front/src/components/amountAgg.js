import React,{Component} from 'react';
import { BorderBox1 } from '@jiaminghi/data-view-react';
import EChartsReact from 'echarts-for-react';


class AmountAgg extends Component{
    constructor(props){
        super(props);
        this.state={
            amountAggEchart:{
                color: ['#d000d0'],
                tooltip:{
                    trigger: 'axis',
                    axisPointer:{
                        type:'cross'
                    }
                },
                grid:{
                    right:'20%'
                },
                toolbox:{
                    feature:{
                        dataView:{show:true,readOnly:false},
                        restore:{show:true},
                        saveAsImage:{show:true}
                    }
                },
                legend:{
                    data:["交易金额"]
                },
                xAxis:[
                    {
                        type:'category',
                        axisTick:{
                            alignWithLabel:true
                        },
                        data:["tody"],
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
                        type:'value',
                        name:'金额',
                        position:"left",
                        alignTicks:true,
                        axisLine:{
                            show:true,
                            lineStyle:{
                                color:['#d000d0']
                            }
                        },
                        nameTextStyle:{
                            color:'#c3dbff',
                            fontSize:12
                        },
                        axisLabel:{
                            formatter: '{value}$',
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
                        name:"金额",
                        type:'line',
                        data:[1]
                    }
                ]
        }
    }
}
    componentDidMount(){

    }
    render(){
        return(
            <div className='xpanel-wrapper xpanel-wrapper-4' style={{position:'relative'}}>
            <div  className='content_title'>交易金额总计</div>
            <BorderBox1>
                <div className='xpanel'>
                    <div className='fill-h'>
                        <EChartsReact  style={{width:'100%',height:'100%'}} option={this.state.amountAggEchart}></EChartsReact>
                    </div>
                </div>
            </BorderBox1>
          </div>
        )
    }

}

export default AmountAgg