import React from 'react';
import { BorderBox1, BorderBox12 ,BorderBox6,BorderBox7} from '@jiaminghi/data-view-react';

const Left = () => {
    return (
        <div className=" fill-h" style={{ width: '25%' }}>
            {/* 面板1 */}
            <div className="xpanel-wrapper xpanel-wrapper-6" style={{ position: 'relative' }}>
                <div className="content_title">xx</div>
                {/* 标题 */}
                <BorderBox12 >
                    <div className="xpanel">
                        <div className="fill-h" id="provinceMap"></div>
                    </div>
                </BorderBox12>
            </div>

            {/* 面板2 */}
            <div
                className="xpanel-wrapper xpanel-wrapper-6"
                style={{ position: 'relative' }}
            >
                <div className="content_title">xx</div>
                {/* 标题 */}
                <BorderBox6>
                    <div className="xpanel">
                        <div className="fill-h" id="cityMap"></div>
                    </div>
                </BorderBox6>
            </div>

            {/* 面板3 */}
            <div
                className="xpanel-wrapper xpanel-wrapper-4"
                style={{ position: 'relative' }}
            >
                <div className="content_title">xx</div>
                {/* 标题 */}
                <BorderBox7>
                    <div className="xpanel">
                        <div className="fill-h" id="countyMap"></div>
                    </div>
                </BorderBox7>
            </div>
        </div>
    );
};

export default Left;
