
import React,{Component} from 'react'

import { BorderBox8 } from '@jiaminghi/data-view-react';

class GasState extends Component{
    constructor(props){
        super(props);
    }

    render(){
        return(
            <div style={{width:"50%",paddingLeft:8}}>
                <BorderBox8 dur={8}>
                    <div className='xpanel'>
                        <div className='fill-h' id="mainMap3"></div>
                    </div>
                </BorderBox8>
            </div>
        )
    }
}

export default GasState;