$( function() {
    $( "#maxCommute" ).slider({
        range: true,
        min: 0,
        max: 180,
        values: [ 15, 120 ],
        slide: function( event, ui ) {
            $( "#amountMaxCommute" ).val(ui.values[ 0 ] + " minutes to " + ui.values[ 1 ] + " minutes." );
        }
    });
    $( "#amountMaxCommute" ).val($( "#maxCommute" ).slider( "values", 0 ) +
        " minutes to " + $( "#maxCommute" ).slider( "values", 1 ) + " minutes" );
} );