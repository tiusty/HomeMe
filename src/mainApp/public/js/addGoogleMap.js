function initMap() {
	var mapDiv = document.getElementById('map');
	var map = new google.maps.Map(mapDiv, {
		center: {lat:42.3601, lng: -71.0589},
		zoom: 8
	});
	distanceMatrix();
	var geocoder = new google.maps.Geocoder();
	geocodeAddress(geocoder, map, myStreetAddress);
	geocodeAddress(geocoder, map, myDestination);
}
