import { BorderBox1 } from '@jiaminghi/data-view-react';
import React,{Component} from 'react'
import ReactECharts from 'echarts-for-react';


const invalidTransactionType=[{"value":0,"name":"合约调用"},{"value":0,"name":"转账"}];
//TODO the pie cant extends the parent windows
class TransactionType extends Component{
    constructor(props){
        super(props);
        this.state={
                transTypeEcharts:{       
                        color:['#9702fe', '#ff893b'],
                        tooltip:{
                            trigger:'item',
                            formatter: '{a} <br/>{b}: {c}({d}%)'
                        },
                        legend:{
                            orient:'vertical',
                            top:20,
                            right:'20%',
                            data:["转账",'合约调用'],
                            textStyle:{
                                fontSize:12,
                                color:'#ffffff'
                            },
                            icon:'circle',
                            itemWidth:10,
                            itemHeight:10,
                            itemGap:10
                        },
                        series:[
                            {
                                name:'交易类型',
                                type:'pie',
                                redius:['50%','70%'],
                                center: ['35%','50%'],
                                avoidLabelOverlap:false,
                                label:{
                                    show:false,
                                    position:'center'
                                },
                                emphasis:{
                                    label:{
                                        show:true,
                                        fontSize:'20',
                                        fontWeight:'bold'
                                    }
                                },
                                labelLine:{
                                    show:false
                                },
                                data:this.asyncGetTransactionType()
                            }
                        ]
                    }
                }
    }
    componentDidMount(){
        // this.setState({
        //     transTypeEcharts:{
        //             color:['#9702fe', '#ff893b'],
        //             tooltip:{
        //                 trigger:'item',
        //                 formatter: '{a} <br/>{b}: {c}({d}%)'
        //             },
        //             legend:{
        //                 orient:'vertical',
        //                 top:30,
        //                 right:'20%',
        //                 data:["转账",'合约调用'],
        //                 textStyle:{
        //                     fontSize:12,
        //                     color:'#ffffff'
        //                 },
        //                 icon:'circle',
        //                 itemWidth:10,
        //                 itemHeight:10,
        //                 itemGap:10
        //             },
        //             series:[
        //                 {
        //                     name:'交易类型',
        //                     type:'pie',
        //                     redius:['50%','70%'],
        //                     center: ['35%','50%'],
        //                     avoidLabelOverlap:false,
        //                     label:{
        //                         show:false,
        //                         position:'center'
        //                     },
        //                     emphasis:{
        //                         label:{
        //                             show:true,
        //                             fontSize:'30',
        //                             fontWeight:'bold'
        //                         }
        //                     },
        //                     labelLine:{
        //                         show:false
        //                     },
        //                     data:this.asyncGetTransactionType()
        //                 }
        //             ]
        //         }
        //     })
    }
    asyncGetTransactionType(){
        let result =undefined;
        try{
            throw("invalid request");
        }catch(error){
            console.log("component/transactionType.js err:",error);
            result=invalidTransactionType;
        }finally{
            return result;
        }
    }
    render(){
       return(
        <div className='xpanel-wrapper xpanel-wrapper-6' style={{position:'relative'}}>
        <div className='content_title'>交易类型占比</div>
        <BorderBox1>
            <div className='xpanel'>
                <div className='fill-h' >
                    <ReactECharts style={{width:'100%',height:'100%'}}  option={this.state.transTypeEcharts}></ReactECharts>
                </div>
            </div>
        </BorderBox1>
    </div>
       )

    }
}


export default TransactionType;