/**
 * To Do
 * - If can't set tides than hide box.
 *
 */
google.charts.load('current', {'packages':['corechart']});

var needs_resizing = true;
var main_images = [
    {
        'image': 'images/services/services_mc_02_1500x1000.jpg',
        'caption': 'Concordia Boatyard',
        'credit': 'Photograph by <a href="https://www.cohenphotography.com">Matthew Cohen Photography</a>',
        'position': 'bottom center'
    },
    {
        'image': 'images/services/services_mc_03_1500x1000.jpg',
        'caption': 'Concordia Engine Room',
        'credit': 'Photograph by <a href="https://www.cohenphotography.com">Matthew Cohen Photography</a>',
        'position': 'bottom center'
    },
    {
        'image': 'images/services/services_mc_04_1500x1000.jpg',
        'caption': 'Concordia Boatyard',
        'credit': 'Photograph by <a href="https://www.cohenphotography.com">Matthew Cohen Photography</a>',
        'position': 'bottom center'
    },
    {
        'image': 'images/services/services_10_1600x900.jpg',
        'caption': 'Concordia Boatyard',
        'credit': '',
        'position': 'bottom center'
    },
    {
        'image': 'images/services/crew_2014_01_1200x856.jpg',
        'caption': '<span class="font_dark">Concordia crew 2014</span>',
        'credit': '<span class="font_dark">Photograph by Carol Hill</span>',
        'position': 'bottom center'
    },
    {
        'image': 'images/yawls/001/large_001_005_1200x800.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=1">#1, Java</a>',
        'credit': 'Photograph by Vagn & Sally Worm',
        'position': 'top left'
    },
    {
        'image': 'images/yawls/014/large_014_007_1280x691.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=14">#14, Saxon</a>',
        'credit': 'Photograph by Dan McGrath',
        'position': 'top left'
    },
    {
        'image': 'images/yawls/015/large_015_013_1023x678.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=15">#15, Lotus</a>',
        'credit': 'Photograph by Bruce Baycroft',
        'position': 'bottom right'
    },
    {
        'image': 'images/yawls/017/large_017_030_1024x680.jpg',
        'caption': 'Concordia sloop <a href="/yawl.php?id=17">#17, Actaea</a>',
        'credit':  null,
        'position': 'center center'
    },
    {
        'image': 'images/yawls/020/large_020_018_1200x778.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=20">#20, Fleetwood</a>',
        'credit': null,
        'position': 'center center'
    },
    {
        'image': 'images/yawls/034/large_034_006_1024x768.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=34">#34, Mandala</a>',
        'credit': 'Photograph by Tony Lincoln',
        'position': 'center center'
    },
    {
        'image': 'images/yawls/044/large_044_010_1500x1000.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=44">#44, Lacerta</a>',
        'credit': 'Photograph by <a href="http://jennifercaseyphotography.com/">Jennifer Casey Photography</a>',
        'position': 'center right'
    },
    {
        'image': 'images/yawls/057/large_057_005_1400x933.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=57">#57, Javelin</a>',
        'credit': 'Photograph by J Michael Stranz',
        'position': 'center center'
    },
    {
        'image': 'images/yawls/058/large_058_009_1200x900.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=58">#58, Off Call</a>',
        'credit': 'Photograph by Chase Castner',
        'position': 'center center'
    },
    {
        'image': 'images/yawls/076/large_076_013_1200x900.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=76">#76, Sumatra</a>',
        'credit': 'Photograph by Doug Cole',
        'position': 'center center'
    },
    {
        'image': 'images/yawls/085/large_085_022_1200x595.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=85">#85, Arapaho</a>',
        'credit': 'Photograph by Carl Cramer',
        'position': 'center center'
    },
    {
        'image': 'images/yawls/092/large_092_002_1024x841.jpg',
        'caption': 'Concordia yawl <a href="/yawl.php?id=92">#92, Eagle</a>',
        'credit': 'Photograph by Steve Baker',
        'position': 'bottom left'
    }

];

$(function(){
    //$("#menu").hide();
    update_image_size();
    $("#menu_icon").on("click", function() {
        $("#menu").toggleClass('hide');
    });

    // Change nav icon/title depending where you are on the page.
    $(window).on("scroll", function(e) {
        if ($(window).scrollTop() > $(window).height()) {
            $("#nav_left").html('<a href="/"><img src="images/concordia_logo_white_50x50.png" class="nav_logo_small"/></a>');
        } else {
            $("#nav_left").html('<a href="/"><img src="images/concordia_logo_white_50x50.png" class="nav_logo_small"/>Concordia Company</a>');
        }
    });

    // Update image size at max every 100ms
    setInterval(update_image_size, 100);
    $(window).resize(function() {
        needs_resizing = true;
    });

    update_yawls();

    update_main_image();

    google.charts.setOnLoadCallback(update_tides);
});


function update_yawls()
{
    for (var i=1; i <= 103; i++) {
        $("#yawl_boxes").append('<div class="box"><a href="yawl.php?id=' + i + '">' + i + '</a></div>');
    }
}

function update_tides()
{
    url = "https://www.concordiaboats.com/get_tides.php";
    $.ajax({
        url: url,
        cache: false,
        dataType: "json"
    }).done(function ( data ) {
        var tides = data['tides'];
        var tide_str = "";

        var google_data = new google.visualization.DataTable();
        google_data.addColumn('datetime', 'Time');
        google_data.addColumn('number', 'Depth');
        google_data.addColumn({type:'string', role:'tooltip'});
        google_data.addColumn({type:'string', role:'annotation'});
        google_data.addColumn({type:'string', role:'annotationText'});

        for (var i in tides) {
            var tide = tides[i];
            tide_str += tide['long_desc'] + "<br/>";
            google_data.addRow([new Date(tide['ts'] * 1000),
                                Math.round(tide['pred'] * 1000),
                                tide['long_desc'],
                                tide['short_desc'],
                                tide['long_desc']]);
        }
        var options = {
            curveType: 'function',
            lineWidth: 3,
            colors: ['#2266bb'],
            legend: { position: 'none' },
            backgroundColor: { fill:'transparent' },
            chartArea: {'width': '80%', 'height': '100%'},
            annotations: {
                textStyle: {
                    bold: true,
                    color: '#000066'
                }
            },
            hAxis: { textPosition: 'none',  gridlines: { color: 'transparent' },  baselineColor: 'transparent' },
            vAxis: { textPosition: 'none',  gridlines: { color: 'transparent' },  baselineColor: 'transparent' }
        };

        var chart = new google.visualization.LineChart(document.getElementById('chart_tides'));
        chart.draw(google_data, options);
    });
}


function update_image_size()
{
    if (needs_resizing == false) {
        return;
    }
    needs_resizing = false;
    $("#main_image").css("height", $(window).height());
}

function update_main_image()
{
    var index = Math.floor(Math.random() * main_images.length);
    var main_image = main_images[index];

    $("#main_image").css('background-image', "url('" + main_image['image'] + "')");
    $("#main_image").css('background-position', main_image['position']);

    if (main_image['credit'] == null) {
        $("#main_title").html(main_image['caption']);
    } else {
        $("#main_title").html(main_image['caption'] + '<br><div class="credit">' + main_image['credit'] + '</div>');
    }
}
