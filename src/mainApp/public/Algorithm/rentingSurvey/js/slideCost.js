$( function() {
    $( "#cost" ).slider({
        range: true,
        min: 0,
        max: 5000,
        values: [ 700, 1200 ],
        slide: function( event, ui ) {
            $( "#amountCost" ).val( "$" + ui.values[ 0 ] + " - $" + ui.values[ 1 ] );
        }
    });
    $( "#amountCost" ).val( "$" + $( "#cost" ).slider( "values", 0 ) +
        " - $" + $( "#cost" ).slider( "values", 1 ) );
} );