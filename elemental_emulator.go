package main

import (
	"fmt"
	"net/http"
)

const xmlString = `
<live_event_list>
	<live_event href="/live_events/3">
		<input>
			<device_input>
				<channel>1</channel>
				<input_format>Auto</input_format>
				<device_type>AJA</device_type>
				<device_number>0</device_number>
			</device_input>
		</input>
		<node_id>2</node_id>
		<notification>
			<email></email>
			<web_callback_url></web_callback_url>
		</notification>
		<user_data>
			<franchise>StarCraft II</franchise>
		</user_data>
		<submitted>2010-01-26 02:07:08 UTC</submitted>
		<status>complete</status>
	</live_event>
	<live_event href="/live_events/4">
		<input>
			<device_input>
				<channel>1</channel>
				<input_format>Auto</input_format>
				<device_type>AJA</device_type>
				<device_number>0</device_number>
			</device_input>
		</input>
		<node_id>2</node_id>
		<notification>
			<email></email>
			<web_callback_url></web_callback_url>
		</notification>
		<user_data>
			<franchise>Shootmania</franchise>
		</user_data>
		<submitted>2010-01-26 02:07:08 UTC</submitted>
		<status>complete</status>
	</live_event>
	<live_event href="/live_events/5">
		<input>
			<device_input>
				<channel>1</channel>
				<input_format>Auto</input_format>
				<device_type>AJA</device_type>
				<device_number>0</device_number>
			</device_input>
		</input>
		<node_id>2</node_id>
		<notification>
			<email></email>
			<web_callback_url></web_callback_url>
		</notification>
		<user_data>
			<franchise>League-of-Legends</franchise>
		</user_data>
		<submitted>2010-01-26 02:07:08 UTC</submitted>
		<status>complete</status>
	</live_event>
</live_event_list>`

func main() {
	http.HandleFunc("/live_events", liveEventList)

	err := http.ListenAndServe(":8009", nil)
	if err != nil {
		fmt.Println("Failed to start elemental emulator server on :8009")
	}
}

func liveEventList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/xml")
	fmt.Fprint(w, xmlString)
}
