$(function() {
    var equipChart = echarts.init( document.getElementById( "chart-tags" ) ),
        trendChart = echarts.init( document.getElementById( "chart-trend" ) );
    window.onresize = function() {
        equipChart.resize();
        trendChart.resize();
    };



    // updateTagsChart( scaleData );

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    var assets = { },
        interval = 211000,
        timer = null;

    // // 날짜
    // $( ".datetime" ).datetimepicker({
    //     format: "yyyy-mm-dd",
    //     pickerPosition : "bottom-left",
    //     todayHighlight : 1,
    //     minView: 2,
    //     maxView: 4,
    //     autoclose: true
    // });
    //
    initializeAssets();
    updateStats();
    // startTimer();
    // // $( "#modal-ipaslog" ).modal( "show" );

    /**
     * 2. 이벤트
     *
     */

    // 자산 선택
    $( "#select-assets" ).on( "change", function() {
        stopTimer();
        updateStats();
        startTimer();
    });


    // $( ".btn-start" ).click(function(e) {
    //     // e.preventDefault();
    //     if ( timer === null ) {
    //         startTimer();
    //     } else {
    //         stopTimer();
    //     }
    // });
    //
    // $( ".activity" ).change(function() {
    //     var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
    //         orgId = asset[0],
    //         groupId = asset[1];
    //     updateLogs( orgId, groupId );
    // });

        // $( ".btn-test" ).click(function(e) {
        //     updateStats();
        // });



    /**
     * 3. 함수
     *
     */

    // 자산 초기화
    function initializeAssets() {
        $.ajax({
            type  : "GET",
            async : true,
            url   : "/userassetclass/1/children"
        }).done( function( result ) {
            // 기관
            $.each( result, function( idx, org ) {
                $( "#select-assets" ).append(
                    $( "<option>", {
                        value: org.asset_id + "/-1",
                        text: org.name
                    })
                );

                // 그룹
                $.each( org.children, function( i, group ) {
                    $( "#select-assets" ).append(
                        $( "<option>", {
                            value: org.asset_id + "/" + group.asset_id,
                            text: "- " + group.name,
                            class: "ml20"
                        })
                    );
                });
            });

        }).always( function() {
            $( "#select-assets" ).selectpicker( "refresh" );
        });
    }

    function startTimer() {  // use a one-off timer
        $( ".btn-start" ).removeClass( "default" ).addClass( "blue" );
        $( ".btn-start .text" ).html( "<i class='fa fa-circle-o-notch fa-spin'></i>" );
        timer = setInterval(updateStats, interval);
    }

    function stopTimer() {
        $( ".btn-start" ).removeClass( "blue" ).addClass( "default" );
        $( ".btn-start .text" ).html( "<i class='fa fa-play'></i>" );
        clearTimeout( timer );
        timer = null;
    }

    function updateStats() {
        var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
            orgId = asset[0],
            groupId = asset[1];

        updateSummary( orgId, groupId );

    //     updateRankings( orgId, groupId );
    //     updateLogs( orgId, groupId );
    //
    //     $( ".text-updated" ).removeClass( "hide" );
    //     setTimeout(function(){ $( ".text-updated" ).addClass( "hide" ); }, 500);
    }


    function updateSummary( orgId, groupId ) {
        var url = "/stats/summary/org/" + orgId + "/group/" + groupId;
        $.ajax({
            type  : "GET",
            async : true,
            url   : url
        }).done( function( r ) {
            // console.log(r);

            updateEquipChart( r.equipCountByType );
            
            // 이벤트 타입 통계
            $( ".count-startup" ).text( r.eventTypes[StartupEvent] );
            $( ".count-shock" ).text( r.eventTypes[ShockEvent] );
            $( ".count-speeding" ).text( r.eventTypes[SpeedingEvent] );
            $( ".count-proximity" ).text( r.eventTypes[ProximityEvent] );

            // 자산 통계
            $( ".count-pt" ).text( r.equipCountByType[PT] );
            $( ".count-zt" ).text( r.equipCountByType[ZT] );
            $( ".count-vt" ).text( r.equipCountByType[VT] );
            $( ".count-total-equips" ).text( r.equipCountByType[PT] + r.equipCountByType[ZT] + r.equipCountByType[VT] );

            //

            //
            var activated = 0;
            $.each(r.activated, function(i, r) {
                activated += r.count;
            });
            $( ".count-activated" ).text( activated );
        }).always( function() {
        });


        updateTrendChart( orgId, groupId );
    }


    function updateTrendChart( orgId, groupId ) {
        // var itemStyle = {
        //     normal: {
        //     },
        //     emphasis: {
        //         barBorderWidth: 1,
        //         shadowBlur: 10,
        //         shadowOffsetX:0,
        //         shadowOffsetY: 0,
        //         shadowColor: 'rgba(0,0,0,0.5)'
        //     }
        // };


        var url = "/stats/timeline/org/" + orgId + "/group/" + groupId;
        $.ajax({
            type  : "GET",
            async : true,
            url   : url
        }).done( function( r ) {
            // console.log(r);
            option = {
                useUTC: true,
                textStyle: {
                    // color: "#ccc",
                    // fontSize: 10
                },
                grid: {
                    top:    50,
                    bottom: 30,
                    left:   '5%',
                    right:  '5%',
                },
                // color: ['#e7505a','#3598dc', '#32c5d2', '#f7ca18', '#8e44ad',           '#749f83',  '#ca8622', '#bda29a','#6e7074', '#546570', '#c4ccd3'],
                color: ['#e35b5a', '#4b77be','#2ab4c0', '#f7ca18', '#8e44ad',           '#749f83',  '#ca8622', '#bda29a','#6e7074', '#546570', '#c4ccd3'],
                // backgroundColor: '#eee',
                legend: {
                    // data: ['Shock', 'Speeding', 'Proximity'],
                    top: 5,
                    textStyle: {
                        // color: "#ccc",
                        // fontSize: 10
                    },
                    align: 'left',
                    left: 10
                },
                tooltip: {
                    trigger: 'axis',
                    // axisPointer: {
                    //     type: 'cross'
                    // },
                    // backgroundColor: 'rgba(245, 245, 245, 0.8)',
                    borderWidth: 1,
                    borderColor: '#777',
                    padding: 5,
                    // textStyle: {
                    //     color: '#000'
                    // },
                },
                xAxis: {
                    // data: xAxisData,
                    boundaryGap : true,
                    // interval: 3600 * 1000,
                    // name: 'X Axis',
                    type : 'time',
                    // silent: true,
                    // offset: 100,
                    axisLine: {onZero: false},
                    splitLine: {show: false},
                    splitArea: {show: false},
                    axisLabel: {
                        showMinLabel: false,
                        showMaxLabel: false,
                        formatter: function(value, index) {
                            var d = moment( value ).utc();
                            // console.log(d.format("MMM D, HH"));
                            return d.format("MMM D, HH");
                            // console.log(value);
                            // return value;
                            // return 3;
                        },
                        rich: {
                            table: {
                                lineHeight: 20,
                                align: 'center'
                            }
                        }
                    }
                },
                yAxis: {
                    splitLine: {
                        show: true,
                        color: '#777',
                        lineStyle: {
                            type: "dotted"
                        }
                    },
                    // splitArea: {show: false},
                    // axisTick: {
                        // Interval: 11

                    // }
                    // inverse: true,
                    // splitArea: {show: false}
                },
                // grid: {
//        left: 100
//        containLabel: true
//                 },
//    visualMap: {
//        type: 'continuous',
//        dimension: 1,
//        text: ['High', 'Low'],
//        inverse: true,
//        itemHeight: 200,
//        calculable: true,
//        min: -2,
//        max: 6,
//        top: 60,
//        left: 10,
//        inRange: {
//            colorLightness: [0.4, 0.8]
//        },
//        outOfRange: {
//            color: '#bbb'
//        },
//        controller: {
//            inRange: {
//                color: '#2f4554'
//            }
//        }
//    },
                series: [
                    {
                        name: 'Shock',
                        type: 'bar',
                        stack: 'event',
                        barMaxWidth: '30px',
                        // itemStyle: itemStyle,
                        data: r.shock,
                    },
                    {
                        name: 'Speeding',
                        type: 'bar',
                        stack: 'event',
                        // itemStyle: itemStyle,
                        data: r.speeding
                    },
                    {
                        name: 'Proximity',
                        type: 'bar',
                        stack: 'event',
                        // itemStyle: itemStyle,
                        data: r.proximity
                    }
                ]
            };

            trendChart.setOption(option, true);
        }).always( function() {
        });

    }


    function updateEquipChart( rawData ) {
        var scaleData = [];
        $.each( rawData, function( tagType, v ) {
            var name;
            if ( tagType == PT ) {
                name = "PT";
            } else if ( tagType == ZT ) {
                name = "ZT";
            } else if ( tagType == VT ) {
                name = "VT";
            } else {
                name = "Unknown";
            }
            scaleData.push( { name: name, value: v } );
        });
        // console.log( scaleData );

        var data = [];
        for (var i = 0; i < scaleData.length; i++) {
            data.push({
                value: scaleData[i].value,
                name: scaleData[i].name,
                itemStyle: {
                    normal: {
                        borderWidth: 8,
                        shadowBlur: 30,
                        borderColor: new echarts.graphic.LinearGradient(0, 0, 1, 1, [{
                            offset: 0,
                            color: '#7777eb'
                        }, {
                            offset: 1,
                            color: '#70ffac'
                        }]),
                        shadowColor: 'rgba(142, 152, 241, 0.6)'
                    }
                }
            }, {
                value: 4,
                name: '',
                itemStyle: {
                    normal: {
                        label: {
                            show: false
                        },
                        labelLine: {
                            show: false
                        },
                        color: 'rgba(0, 0, 0, 0)',
                        borderColor: 'rgba(0, 0, 0, 0)',
                        borderWidth: 0
                    }
                }
            });
        }
        var seriesObj = [{
            name: '',
            type: 'pie',
            clockWise: false,
            radius: [95, 100],
            hoverAnimation: false,
            itemStyle: {
                normal: {
                    label: {
                        show: true,
                        position: 'outside',
                        color: '#333',
                        fontSize:18,
                        fontWeight: "bold",
                        formatter: function (params) {
                            // var percent = 0;
                            var total = 0;
                            for (var i = 0; i < scaleData.length; i++) {
                                total += scaleData[i].value;
                            }
                            var percent = ((params.value / total) * 100).toFixed(0);
                            if (params.name !== '') {
                                return params.name + '\n{white|' + '' + params.value + '}';
                            } else {
                                return '';
                            }
                        },
                        rich: {
                            white: {
                                color: '#333',
                                align: 'center',
                                padding: [5, 0]
                            }
                        }
                    },
                    labelLine: {
                        show: false
                    }
                }
            },
            data: data
        }];

        // var rich = {
        //     white: {
        //         color: '#ddd',
        //         align: 'center',
        //         padding: [5, 0]
        //     }
        // };
        // var placeHolderStyle = {
        //     normal: {
        //         label: {
        //             show: false
        //         },
        //         labelLine: {
        //             show: false
        //         },
        //         color: 'rgba(0, 0, 0, 0)',
        //         borderColor: 'rgba(0, 0, 0, 0)',
        //         borderWidth: 0
        //     }
        // };

        var option = {
            // backgroundColor: '#04243E',
            tooltip: {
                show: false
            },
            legend: {
                show: false
            },
            toolbox: {
                show: false
            },
            series: seriesObj
        }

        if (option && typeof option === "object") {
            equipChart.setOption(option, true);
        }
    }


    // function updateRankings( orgId, groupId ) {
    //     $( ".table-ranking" ).each(function( idx, obj ) {
    //         var url = $( this ).data( "query" ) + "/org/" + orgId + "/group/" + groupId;
    //         $( this ).bootstrapTable( "refresh", {
    //             url: url,
    //             silent: true
    //         });
    //     });
    // }
    //
    //
    // function updateLogs( orgId, groupId ) {
    //     var url = '/getRealTimeLogs?limit=5';
    //     var activities = [];
    //     $( ".activity" ).each(function( idx, obj ) {
    //         if ( $( obj ).is(":checked") ) {
    //             activities.push( "event_type=" + $( obj ).val() );
    //         }
    //     });
    //
    //     if ( activities.length < 1 ) {
    //         $( "#table-ipaslogs" ).bootstrapTable( "removeAll" );
    //     } else {
    //         var urlSuffix = "";
    //         if ( orgId > 0 ) {
    //             urlSuffix += "&org_id=" + orgId;
    //             if ( groupId > 0 ) {
    //                 urlSuffix += "&group_id=" + groupId;
    //             }
    //         }
    //         $( "#table-ipaslogs" ).bootstrapTable( "refresh", {
    //             url: url + "&" + activities.join( "&" ) + urlSuffix,
    //             silent: true
    //         });
    //     }
    // }

});

