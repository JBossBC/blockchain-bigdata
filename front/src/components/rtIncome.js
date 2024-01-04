import { ScrollRankingBoard ,BorderBox13} from '@jiaminghi/data-view-react';
import React,{Component} from 'react'

const invalidIncome ={data:[{name:"xiyang收入",value:"5tron"}],carousel:'page'};

class RealTimeIncome extends Component{
    constructor(props){
        super(props);
        this.state={
            rtIncome:[],
        }
    };
    componentDidMount(){
        this.setState({rtIncome:this.asyncGetIncome()});
    }
    asyncGetIncome(){
        let result =[];
        try{
            throw("bad request");
        }catch (error){
            result=invalidIncome;
        }finally{
            return result;
        }
    }
    render(){
        return( <div className='xpanel-wrapper xpanel-wrapper-middle'>
            <BorderBox13>
                <div className='xpanel'>
                    <div className='fill-h'>
                        <ScrollRankingBoard config={this.state.rtIncome}></ScrollRankingBoard>
                    </div>
                </div>
            </BorderBox13>
        </div>
        )
    }
}

export default RealTimeIncome;