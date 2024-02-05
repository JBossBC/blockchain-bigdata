import React from 'react';
import { BorderBox1,BorderBox8,BorderBox13,BorderBox2 } from '@jiaminghi/data-view-react';

const Right = () => {
    return (
        <div className="col-lg-3 fill-h" style={{ width: '25%' }}>
            {/* 面板1 */}
            <div
                className="xpanel-wrapper xpanel-wrapper-6"
                style={{ position: 'relative' }}
            >
                <div className="content_title">xx</div>
                {/* 标题 */}
                <BorderBox8>
                    <div className="xpanel">
                        <div className="fill-h" id="provinceMap"></div>
                    </div>
                </BorderBox8>
            </div>

            {/* 面板2 */}
            <div
                className="xpanel-wrapper xpanel-wrapper-6"
                style={{ position: 'relative' }}
            >
                <div className="content_title">hh</div>
                {/* 标题 */}
                <BorderBox13>
                    <div className="xpanel">
                        <div className="fill-h" id="cityMap"></div>
                    </div>
                </BorderBox13>
            </div>

            {/* 面板3 */}
            <div
                className="xpanel-wrapper xpanel-wrapper-4"
                style={{ position: 'relative' }}
            >
                <div className="content_title">hh</div>
                {/* 标题 */}
                <BorderBox2>
                    <div className="xpanel">
                        <div className="fill-h" id="countyMap"></div>
                    </div>
                </BorderBox2>
            </div>
        </div>
    );
};

export default Right;
