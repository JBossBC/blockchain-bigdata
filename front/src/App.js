/*
import React,{Component} from 'react'
import './App.css';
import { AccountScreen, AmountAgg, GasState, RealTimeIncome, RealTimePayOff, TokenType, TransactionState, TransactionType } from './components';


class App extends Component{
   constructor(props){
    super(props);

   }
   render(){
    // const {topData,tableData}=this.state;
    return(
      <div className='data'>
        <header className='header_main'>
          <div className='left_bg'></div>
          <div className='right_bg'></div>
          <h3>区块链资金分析平台</h3>
        </header>
        <div className='wrapper'>
        <div className='container-fluid'>
          <div className='row fill-h' style={{display:'flex'}}>
            <div className='col-lg-3 fill-h' style={{width:"25%"}}>
              <RealTimeIncome/>
              <RealTimePayOff/>
            </div>
            <div className='col-lg-6 fill-h' style={{width:"50%"}}>
             <AccountScreen/>
              <div className='xpanel-wrapper xpanel-wrapper-4' style={{display:'flex'}}>
               <TransactionState/>
               <GasState/>
              </div>
            </div>
            <div className='col-lg-3 fill-h' style={{width:"25%"}}>
                  <TransactionType/>
                <TokenType/>
                <AmountAgg/>
            </div>
          </div>
        </div>
        </div>
      </div>
    )
   }
}

export default App;*/
import React, { Component } from 'react';
import {Header,Center,Left,Right} from './components'

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    render() {
        return (
            <div className="data">
                <Header />
                <div className="wrapper">
                    <div className="container-fluid">
                        <div className="row fill-h" style={{ display: 'flex' }}>

                            <Left />
                            <Center />
                            <Right />

                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default App;
