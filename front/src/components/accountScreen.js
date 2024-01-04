
import React,{Component} from 'react'
import EChartsReact from 'echarts-for-react';
import  {BorderBox10, Decoration1} from '@jiaminghi/data-view-react'
// import echarts from 'echarts';
import {NumberToByteSlice} from '../utils/integer';
const upColor = '#00da3c';
const downColor = '#ec0000';
class AccountScreen extends Component{
    constructor(props){
        super(props);
        this.state={
          accountSum:0,
          //价值变化.... MainMap
          pricesBrushEcharts:{
             animation:false,
             legend:{
              bottom:10,
              left:'center',
              data:[]
             },
             tooltip:{
              trigger:'axis',
              axisPointer:{
                type:'cross'
              },
              borderWidth:1,
              borderColor:'#ccc',
              padding:10,
              textStyle:{
                color:'#000'
              },
              position:function(pos,params,el,elReact,size){
                const obj = {
                  top:10
                };
                obj[['left','right'][+(pos[0]<size.viewSize[0/2])]]=30;
                return obj;
              }
             },
             axisPointer:{
              link:[
                {
                  xAxisIndex:'all'
                }
              ],
              label:{
                backgroundColor:'#777'
              }
             },
             toolbox:{
              feature:{
                dataZoom:{
                  yAxisIndex:false
                },
                brush:{
                  type:['lineX','clear']
                }
              }
             },
             bursh:{
              xAxisIndex:"all",
              brushLink:'all',
              outOfBrush:{
                colorAlpha:0.1
              }
             },
             visualMap:{
              show:false,
              seriesIndex:5,
              dimension:2,
              pieces:[
                {
                  value:1,
                  color:downColor
                },{
                  value:-1,
                  color:upColor
                }
              ]
             },
             grid:[
              {
                left:'10%',
                right:'8%',
                height:'50%'
              },
              {
                left:'10%',
                right:'8%',
                top:'63%',
                height:'16%'
              }
             ],
             xAxis:[
              {
                type:'category',
                data:[],
                boundaryGap:false,
                axisLine:{
                  onZero:false
                },
                splitLine:{
                  show:false
                },
                min:'dataMin',
                max:'dataMax',
                axisPointer:{
                  z:100
                }
              },
              {
                type:'category',
                gridIndex:1,
                data:[],
                boundaryGap:false,
                axisLine:{
                  onZero:false
                },
                axisTick:{
                  show:false
                },
                splitLine:{
                  show:false
                },
                axisLabel:{
                  show:false
                },
                min:'dataMin',
                max:'dataMax'
              }
             ],
             yAxis: [
              {
                scale: true,
                splitArea: {
                  show: true
                }
              },
              {
                scale: true,
                gridIndex: 1,
                splitNumber: 2,
                axisLabel: { show: false },
                axisLine: { show: false },
                axisTick: { show: false },
                splitLine: { show: false }
              }
            ],
            dataZoom: [
              {
                type: 'inside',
                xAxisIndex: [0, 1],
                start: 98,
                end: 100
              },
              {
                show: true,
                xAxisIndex: [0, 1],
                type: 'slider',
                top: '85%',
                start: 98,
                end: 100
              }
            ],
            series: [
              {
                name: '价格(美元)',
                type: 'candlestick',
                data: [],
                itemStyle: {
                  color: upColor,
                  color0: downColor,
                  borderColor: undefined,
                  borderColor0: undefined
                }
              }
            ]
          }
        };
        // this.setState({accountSum:NumberToByteSlice(null)},()=>{
        //   //axios request
        //   let result = undefined;
        //   this.state.accountSum  =  result;
        // });
       }
    asyncGetPrices(){
      let result=[];
      try{
        
      }catch(error){
        console.log("accountScreen.js error",error);
        result=[];
      }finally{
        return result;
      }
    }   
    componentDidMount(){
      this.setState({accountSum:NumberToByteSlice(100000000).map((v, k) => (<div key={k} id={k} className='databg'>{v}</div>))},()=>{
        //axios request
        // let result = undefined;
        // this.setState({accountSum : result});
      });
    }   
    render(){
        return(
            <div className='xpanel-wrapper xpanel-wrapper-5'>
            <div className='xpanel' style={{padding:'0px',position:"relative",boxSizing:"border-box", maxWidth:"100%",display:'grid',columnGap:'20px',rowGap:'8px',gridTemplateColumns:'3fr 1fr',gridTemplateRows:'1fr 5fr'}}>
              {/* <div className='map_bg'></div> */}
              <div className='circle_allow'></div>
              {/* <div className='circle_bg'></div> */}
              <div style={{top:10,display:"flex",justifyContent:"center",color:'#fff',alignItems:"center"}}>
                <p style={{marginRight:"10px",fontSize:"18px"}}>交易账户总数量:</p>
                {this.state.accountSum}
              </div>
              <div>
                <Decoration1 style={{width:'100%',height:'100%'}}></Decoration1>
              </div>
              {/* TODO this component will expands so that the parent element is smaller than this component */}
              <div className='fill-h' style={{gridColumnStart:1,gridColumnEnd:3}}>
                <BorderBox10 style={{maxWidth:"100%"}}>
                <EChartsReact style={{width:'100%',height:'100%'}} option={this.state.pricesBrushEcharts}></EChartsReact>
                </BorderBox10>
              </div>
            </div>
          </div>
        )
    }
}

export default AccountScreen;