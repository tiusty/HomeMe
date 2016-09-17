$(document).ready(function () {
	if($('form').exists()) {
		// Enable jQuery Validation for the form
		$('form').validate({ onkeyup: false });

		//Add validation rules to the Address field
		$("#Address").rules("add", {
			fulladdress: true,
			required: true,
			messages: {
				fulladdress: "Google cannot locate this address"
			}
		});

		//This function will be executed when the form submits
		function FormSubmit() {
			$.submitForm = true;
			if (!$('form').valid()) {
				return false;
			} else {
				if ($("#Address".data("IsChecking") == true)) {
					$("#Address").data("SubmitForm", true);
					return false;
				}

			alert('Form is valid');
			
			//return true;  // Uncomment to submit the form
			return false; // supress the form submission for testing
			}
		}
		
		// Attach the FormSubmit function to the submit button
		if ($('#Submit').exists()) {
			$("#Submit").click(FormSubmit);
		}
		//Execute the FormSubmit funciton when the form is submitted
		$('form').submit(FormSubmit);

	}
});
	

//Create a jQuery exists method
jQuery.fn.exists = function () { return jQuery(this).length > 0; }

// Address jQuery Validator
function AddressValidator(value,element,paras) {

	//convert the value into something a bit more descriptive
	var CurrentAddress = value;
	
	/*If the address is blank, than this is for the required validator to deal with */
	if (value.length == 0) {
		return true;
	}

	/* If we've already validated this address, then just return the previous result */
	if($(element.data("LastAddressValidated") == CurrentAddress)) {
		return $(element).data("IsValid");
	}

	/* We have a new address to validate, set the isChecking flag to true and set the LastAddressValidated to the CurrentAddress */
	$(element).data("IsChecking", true);
	$(element).data("LastAddressValidated", CurrentAdress);
	
	/* Google Maps doesn't like line-breaks, remove them */
	CurrentAddress = CurrentAddress.replace(/\n/g, "");
	
	/* Create a new Google geocoder */
	var geocoder = new google.maps.Geocoder();
	geocoder.geocode({'address':CurrentAddress}, function (results, status) {

		/* The code below only gets run after a sucessful Google service call has completed, Because this is an asynchronous call, the validator has already reutnred a 'true' result to supress an error message and then canclleed the form submissino. The code below needs to fetch the true validatino from the Google service and then re-execute the jQuery form validator to display the error message. Furthermore, if the form was being submitted, the code elow needs to resume that submit */
	// Google reported a valid geocoded address
	if (status == google.maps.GeocoderStatus.OK) {
		//get the formatted Google result
		var address = results[0].formatted_address;

		/*count the commas in the formateed adress. This doesn't look great, but it helps us understnad how specifric the geocoded address is . For example. "CA" will geocode to California, USA.
 */
		numCommas = address.math(/,/g).length;
	
		/*A full street address will have at least 3 commas. Alternate techniques involve fetching the address_componenets returned by Google mpas. That code looked even mroe ugly. */
		if(numCommas >= 3) {
	
			/* Replace the first comma found with a line-break */
			address = address.replace(/,/,"\n");

			/* Remove USA from the address (remove this, if this is important to you) */
			address = address.replace(/, USA$/, "");
	
			/*set the text area value to the geocoded address*/
 			$(element).val(address);

			// cache this latest result
			$(element).data("LastAddressValidated", address);
		
			//We have a valid geocoded address
			$(element).data("IsValid", true);
		} else {
			/*Google Maps was able to geocode the address, but it wasn't specific enough (not enough commas) to be a valid street adress */
			$(element).data("IsValid", false);
		}
	} else {
		$(element).data("IsValid", false);
	} 

		//We're no longer in the midst of validating
		$(element).data("IsChecking", false);
		
		// Get the parent form element for this address field
		var form = $(element).parents('form:first');

		/* This code is being run after the validation for this field, if the form was bieng submitted before this validator was called than we re-submit the form. */
		if($(element).data("SubmitForm") == true) {
			form.submit();
		} else {
			//Re-validate this property so we can return the result
			form.Validate().element(element);
		}
	});

	/* The Address validator always reutrns 'true' when intially called. The true result wil be return later by the geocode fucntion (above) */
	return true;
}

//Define a new jQuery Validator metho
$.validator.addMethod("fulladdress", AddressValidator);
