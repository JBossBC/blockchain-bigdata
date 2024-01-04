import { BorderBox13, ScrollRankingBoard } from '@jiaminghi/data-view-react';
import React,{Component} from'react'

const invalidPayoff ={data:[{name:"xiyang支出",value:"5tron"}],carousel:'page'};

const payoffEventURL = "";

class RealTimePayOff extends Component{
    constructor(props){
        super(props);
        this.state={
            rtPayOff:[],
            payoffEvent:undefined,
        }
    }
    componentDidMount(){
        this.setState({rtPayOff:this.asyncGetPayoff()});
        this.setState({payoffEvent:new EventSource(payoffEventURL)},()=>{
            this.state.payoffEvent.onmessage=function(event){

            }
        });
    }
    componentWillUnmount(){

    }
    asyncGetPayoff(){
        let result =[];
        try{
            throw("bad request");
        }catch(error){
            console.log("components/rtpayoff.js error",error);
            result=invalidPayoff;
        }finally{
            return result;
        }

    }
    render(){
        return( 
        <div className='xpanel-wrapper xpanel-wrapper-middle'>
            <BorderBox13>
                <div className='xpanel'>
                    <div className='fill-h'>
                        <ScrollRankingBoard config={this.state.rtPayOff}></ScrollRankingBoard>
                    </div>
                </div>
            </BorderBox13>
        </div>
       )
    }
}



export default RealTimePayOff;