$(function() {

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    // var assets = { },
    //     interval = 211000,
    //     timer = null;
    //
    // // // 날짜
    // // $( ".datetime" ).datetimepicker({
    // //     format: "yyyy-mm-dd",
    // //     pickerPosition : "bottom-left",
    // //     todayHighlight : 1,
    // //     minView: 2,
    // //     maxView: 4,
    // //     autoclose: true
    // // });
    //
    // Chart - 이벤트 유형
    var eventTypeChart = Morris.Donut({
        element: 'chart-eventType',
        colors: [ ShockColor, SpeedingColor, ProximityColor ],

        resize: true,
        data: [
            { label: "N/A", value: 0 }
        ],
        formatter: function (x) { return x + ""}
    }).on('click', function(i, row){
        console.log(i, row);
    });

    var trendChart = echarts.init( document.getElementById( "chart-trend" ) );

    initializeAssets();
    updateStats();
    // startTimer();
    //
    /**
     * 2. 이벤트
     *
     */

    // 자산 선택
    $( "#select-assets" ).on( "change", function() {
    //     stopTimer();
        updateStats();
    //     startTimer();
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
    //



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

    // function startTimer() {  // use a one-off timer
    //     $( ".btn-start" ).removeClass( "default" ).addClass( "blue" );
    //     $( ".btn-start .text" ).html( "<i class='fa fa-circle-o-notch fa-spin'></i>" );
    //     timer = setInterval(updateStats, interval);
    // }
    //
    // function stopTimer() {
    //     $( ".btn-start" ).removeClass( "blue" ).addClass( "default" );
    //     $( ".btn-start .text" ).html( "<i class='fa fa-play'></i>" );
    //     clearTimeout( timer );
    //     timer = null;
    // }
    //
    //
    function updateStats() {
        var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
            orgId = asset[0],
            groupId = asset[1];

        updateSummary( orgId, groupId );
        updateEventTags( orgId, groupId );
        updateTrendChart( orgId, groupId );
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
            console.log(r);
            $( ".count-startup" ).text( r.eventTypes[StartupEvent] );
            $( ".count-shock" ).text( r.eventTypes[ShockEvent] );
            $( ".count-speeding" ).text( r.eventTypes[SpeedingEvent] );
            $( ".count-proximity" ).text( r.eventTypes[ProximityEvent] );

            $( ".count-pt" ).text( r.equipCountByType[PT] );
            $( ".count-zt" ).text( r.equipCountByType[ZT] );
            $( ".count-vt" ).text( r.equipCountByType[VT] );
            $( ".count-total-tags" ).text( r.equipCountByType[PT] + r.equipCountByType[ZT] + r.equipCountByType[VT] );
    //
            var  data = [
                {value:  r.eventTypes[ShockEvent], label: 'SHOCK'},
                {value:  r.eventTypes[SpeedingEvent], label: 'SPEEDING'},
                {value:  r.eventTypes[ProximityEvent], label: 'PROXIMITY'},
            ];

            var total = r.eventTypes[ShockEvent] + r.eventTypes[SpeedingEvent] + r.eventTypes[ProximityEvent];
            if ( total === 0 ) {
                total = 1;
            }
            $( "#pgb-shock" ).css( "width", (r.eventTypes[ShockEvent] / total * 100) + "%" );
            $( "#pgb-speeding" ).css( "width", (r.eventTypes[SpeedingEvent] / total * 100) + "%" );
            $( "#pgb-proximity" ).css( "width", (r.eventTypes[ProximityEvent] / total * 100) + "%" );
            // $('#progress-upload .progress-bar').css('width', progress + '%');
            if ( total > 0 ) {
                eventTypeChart.setData( data );
            } else {
                eventTypeChart.setData( [ { value: 0, label: 'N/A' } ] );
            }

            // Update activated equipments
            $( "#table-activated" ).bootstrapTable( "load", r.activated );
        }).always( function() {
        });
    }
    //
    //
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
    function updateEventTags( orgId, groupId ) {
        var url = '/getRealTimeLogs?',
            param = {
                limit: 30
            };
        if ( orgId > 1 ) {
            param.org_id = orgId;
            if ( groupId > 1 ) {
                param.group_id = groupId;
            }
        }
        url += $.param( param );
        // console.log(url);

        var tags = '';
        $.ajax({
            type  : "GET",
            async : true,
            url   : url
        }).done( function( rows ) {
            $.each( rows, function( i, r ) {
                // console.log(r);
                var btnCss,
                    icon;
                if ( r.event_type === StartupEvent ) {
                    icon = "icon-power";
                    btnCss = "default"

                } else if ( r.event_type === ShockEvent ) {
                    icon = "fa fa-bolt";
                    btnCss = "blue";

                } else if ( r.event_type === SpeedingEvent ) {
                    icon = "icon-speedometer";
                    btnCss = "green";

                } else if ( r.event_type === ProximityEvent ) {
                    icon = "icon-size-actual";
                    btnCss = "red-haze";
                } else {
                    icon = "icon-question";
                    btnCss = "yellow-gold"
                }

                // tags += '<div class="label label-' + labelSuffix + ' uppercase" style="display: block; display: inline-block; margin-right: 5px;">';
                // tags += r.equip_id;
                // tags += '</div>';
                tags += '<div class="col-sm-4 mb5"><button type="button" class=" btn ' + btnCss + ' btn-block btn-xs mr5"><i class="' + icon + '"></i> ';
                tags += r.equip_id;
                tags += '</button></div>';
            });

        }).always( function() {
            $( "#event-tags" ).html( tags );
        });
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

});

