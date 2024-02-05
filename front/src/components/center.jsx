import React from 'react';
import { BorderBox1 } from '@jiaminghi/data-view-react';
import { Button } from 'antd';

const Center = () => {
    return (
        <div className="col-lg-6 fill-h" style={{ width: '50%', position: 'relative' }}>
            {/* 面板3 */}
            <div className="xpanel-wrapper xpanel-wrapper-1">
                <BorderBox1>
                    <div className="xpanel" style={{ position: 'relative' }}>
                        <div className="map_bg"></div>
                        <div className="circle_bg"></div>
                        <div className="container" style={{flexDirection:'column'}}>
                            <div>
                                <div className="hexagons position0">
                                    <div className="search"><Button className="buttonk">按钮1</Button></div>
                                </div>

                            </div>
                            <div>
                                <div className="hexagons position1">
                                    <div className="movies"></div>
                                </div>
                                <Button className="button2">按钮2</Button>
                            </div>
                            <div className="hexagons position2">
                                <div className="writing"></div>
                                <Button className="button3">按钮2</Button>
                            </div>
                            <div className="hexagons position3">
                                <div className="police"></div>
                            </div>
                            <div className="hexagons position4">
                                <div className="location"></div>
                            </div>
                            <div className="hexagons position5">
                                <div className="ie"></div>
                            </div>
                        </div>
                    </div>
                </BorderBox1>
            </div>
        </div>
    );
};

export default Center;
