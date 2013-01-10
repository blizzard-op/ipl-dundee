package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const xmlString = `<?xml version="1.0" encoding="UTF-8"?>
<live_event_list>
<live_event href="/live_events/74" version="1.5.4.1235" product="Elemental Live">
  <name>01-08-13 LoL 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>1</device_number>
      <device_id>2</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-09T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-08T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>10</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-07 11:07:08 -0800</submitted>
  <status>pending</status>
  <average_fps></average_fps>
  <buffer_avg></buffer_avg>
  <buffer_max></buffer_max>
  <dropped_frames></dropped_frames>
  <start_time></start_time>
  <elapsed></elapsed>
</live_event>
<live_event href="/live_events/73" version="1.5.4.1235" product="Elemental Live">
  <name>01-08-13 SC2 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>3</device_number>
      <device_id>4</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-09T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-08T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPLWatermarkSC2.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-07 11:07:06 -0800</submitted>
  <status>pending</status>
  <average_fps></average_fps>
  <buffer_avg></buffer_avg>
  <buffer_max></buffer_max>
  <dropped_frames></dropped_frames>
  <start_time></start_time>
  <elapsed></elapsed>
</live_event>
<live_event href="/live_events/72" version="1.5.4.1235" product="Elemental Live">
  <name>01-08-13 SM 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>0</device_number>
      <device_id>1</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-09T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-08T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-07 11:07:04 -0800</submitted>
  <status>pending</status>
  <average_fps></average_fps>
  <buffer_avg></buffer_avg>
  <buffer_max></buffer_max>
  <dropped_frames></dropped_frames>
  <start_time></start_time>
  <elapsed></elapsed>
</live_event>
<live_event href="/live_events/71" version="1.5.4.1235" product="Elemental Live">
  <name>01-08-13 LoL 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>1</device_number>
      <device_id>2</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-08T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-07T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>10</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-06 11:06:59 -0800</submitted>
  <status>running</status>
  <average_fps>29.96</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>14</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-07 09:14:59 -0800</start_time>
  <elapsed>7373</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-07T09:15:02-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/70" version="1.5.4.1235" product="Elemental Live">
  <name>01-07-13 SC2 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>3</device_number>
      <device_id>4</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-08T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-07T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPLWatermarkSC2.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-06 11:06:57 -0800</submitted>
  <status>running</status>
  <average_fps>29.97</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>16</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-07 09:15:01 -0800</start_time>
  <elapsed>7363</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-07T09:15:05-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/69" version="1.5.4.1235" product="Elemental Live">
  <name>01-07-13 SM 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>0</device_number>
      <device_id>5</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-08T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-07T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-06 11:06:55 -0800</submitted>
  <status>running</status>
  <average_fps>29.98</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>11</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-07 09:15:00 -0800</start_time>
  <elapsed>7376</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-07T09:15:02-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/68" version="1.5.4.1235" product="Elemental Live">
  <name>01-06-13 LoL 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>1</device_number>
      <device_id>2</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-07T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-06T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>10</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-05 11:06:50 -0800</submitted>
  <status>complete</status>
  <average_fps>29.93</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>16</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-06 09:14:59 -0800</start_time>
  <complete_time>2013-01-07 09:10:05 -0800</complete_time>
  <elapsed>86096</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-06T09:15:02-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-07T09:10:01-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/67" version="1.5.4.1235" product="Elemental Live">
  <name>01-06-13 SC2 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>3</device_number>
      <device_id>4</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-07T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-06T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPLWatermarkSC2.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-05 11:06:48 -0800</submitted>
  <status>complete</status>
  <average_fps>25.24</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>66</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-06 09:15:01 -0800</start_time>
  <complete_time>2013-01-07 09:10:04 -0800</complete_time>
  <elapsed>86094</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-06T09:15:05-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-07T09:10:01-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/66" version="1.5.4.1235" product="Elemental Live">
  <name>01-06-13 SM 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>0</device_number>
      <device_id>5</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-07T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-06T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-05 11:06:46 -0800</submitted>
  <status>complete</status>
  <average_fps>30.05</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>16</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-06 09:14:59 -0800</start_time>
  <complete_time>2013-01-07 09:10:04 -0800</complete_time>
  <elapsed>86095</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-06T09:15:02-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-07T09:10:01-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/65" version="1.5.4.1235" product="Elemental Live">
  <name>01-05-13 LoL 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>1</device_number>
      <device_id>2</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-06T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-05T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>10</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-04 11:06:40 -0800</submitted>
  <status>complete</status>
  <average_fps>25.62</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>68</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-05 09:14:59 -0800</start_time>
  <complete_time>2013-01-06 09:10:05 -0800</complete_time>
  <elapsed>86097</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-05T09:15:03-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-06T09:10:02-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/64" version="1.5.4.1235" product="Elemental Live">
  <name>01-05-13 SC2 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>3</device_number>
      <device_id>4</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-06T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-05T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPLWatermarkSC2.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-04 11:06:38 -0800</submitted>
  <status>complete</status>
  <average_fps></average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>100</buffer_max>
  <dropped_frames>24</dropped_frames>
  <start_time>2013-01-05 09:15:01 -0800</start_time>
  <complete_time>2013-01-06 09:10:05 -0800</complete_time>
  <elapsed>86095</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-05T09:15:05-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-06T09:10:01-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/63" version="1.5.4.1235" product="Elemental Live">
  <name>01-05-13 SM 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>0</device_number>
      <device_id>5</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-06T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-05T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-04 11:06:36 -0800</submitted>
  <status>complete</status>
  <average_fps>30.0</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>11</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-05 09:14:59 -0800</start_time>
  <complete_time>2013-01-06 09:10:04 -0800</complete_time>
  <elapsed>86100</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-05T09:15:02-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-06T09:10:01-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/62" version="1.5.4.1235" product="Elemental Live">
  <name>01-04-13 LoL 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>1</device_number>
      <device_id>2</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-05T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-04T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>10</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-03 11:06:31 -0800</submitted>
  <status>complete</status>
  <average_fps></average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>79</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-04 09:14:59 -0800</start_time>
  <complete_time>2013-01-05 09:10:04 -0800</complete_time>
  <elapsed>86098</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-04T09:15:02-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-05T09:10:01-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/61" version="1.5.4.1235" product="Elemental Live">
  <name>01-04-13 SC2 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>3</device_number>
      <device_id>4</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-05T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-04T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPLWatermarkSC2.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-03 11:06:29 -0800</submitted>
  <status>complete</status>
  <average_fps>30.0</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>64</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-04 17:51:39 -0800</start_time>
  <complete_time>2013-01-05 09:10:06 -0800</complete_time>
  <elapsed>55100</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-04T17:51:42-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-05T09:10:01-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/60" version="1.5.4.1235" product="Elemental Live">
  <name>01-04-13 SM 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>0</device_number>
      <device_id>5</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-05T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-04T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-03 11:06:27 -0800</submitted>
  <status>complete</status>
  <average_fps>26.6</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>11</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-04 09:14:59 -0800</start_time>
  <complete_time>2013-01-05 09:10:05 -0800</complete_time>
  <elapsed>86098</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-04T09:15:02-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-05T09:10:02-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/59" version="1.5.4.1235" product="Elemental Live">
  <name>01-03-13 LoL 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>1</device_number>
      <device_id>2</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-04T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-03T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>10</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-02 11:06:22 -0800</submitted>
  <status>complete</status>
  <average_fps></average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>18</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-03 09:14:59 -0800</start_time>
  <complete_time>2013-01-04 09:10:05 -0800</complete_time>
  <elapsed>86092</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-03T09:15:07-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-04T09:10:00-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/58" version="1.5.4.1235" product="Elemental Live">
  <name>01-03-13 SC2 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>3</device_number>
      <device_id>4</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-04T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-03T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-02 11:06:20 -0800</submitted>
  <status>complete</status>
  <average_fps>29.96</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>22</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-03 18:38:16 -0800</start_time>
  <complete_time>2013-01-04 09:10:04 -0800</complete_time>
  <elapsed>52295</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-03T18:38:20-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-04T09:10:00-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/57" version="1.5.4.1235" product="Elemental Live">
  <name>01-03-13 SM 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>0</device_number>
      <device_id>5</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-04T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-03T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-02 11:06:17 -0800</submitted>
  <status>complete</status>
  <average_fps>25.28</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>20</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-03 09:15:00 -0800</start_time>
  <complete_time>2013-01-04 09:10:04 -0800</complete_time>
  <elapsed>86098</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-03T09:15:03-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-04T09:10:02-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/56" version="1.5.4.1235" product="Elemental Live">
  <name>01-02-13 LoL 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>1</device_number>
      <device_id>2</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-03T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-02T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>10</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-01 11:06:13 -0800</submitted>
  <status>complete</status>
  <average_fps>30.0</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>18</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-02 09:14:59 -0800</start_time>
  <complete_time>2013-01-03 09:10:03 -0800</complete_time>
  <elapsed>86099</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-02T09:15:03-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-03T09:10:00-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
<live_event href="/live_events/55" version="1.5.4.1235" product="Elemental Live">
  <name>01-02-13 SC2 1</name>
  <input>
    <failover_time>5</failover_time>
    <order>1</order>
    <program_id nil="true"></program_id>
    <device_input>
      <channel>1</channel>
      <input_format>Auto</input_format>
      <device_type>AJA</device_type>
      <device_number>3</device_number>
      <device_id>4</device_id>
    </device_input>
  </input>
  <loop_all_inputs>false</loop_all_inputs>
  <timecode_source>embedded</timecode_source>
  <failure_rule>
    <priority>50</priority>
    <restart_on_failure>false</restart_on_failure>
  </failure_rule>
  <timing>
    <end_at>2013-01-03T09:10:00-08:00</end_at>
    <end_type>end_at</end_type>
    <start_at>2013-01-02T09:15:00-08:00</start_at>
    <start_at_buffer>0</start_at_buffer>
    <start_type>start_at</start_type>
  </timing>
  <notification>
    <email>iplvideotech@ign.com</email>
    <web_callback_url></web_callback_url>
    <on_started>true</on_started>
    <on_complete>true</on_complete>
    <on_alert>true</on_alert>
    <on_warning>true</on_warning>
    <on_error>true</on_error>
  </notification>
  <image_inserter>
    <image_x>0</image_x>
    <image_y>0</image_y>
    <opacity>100</opacity>
    <image_inserter_input>
      <uri>/home/elemental/files/IPL watermark.png</uri>
    </image_inserter_input>
  </image_inserter>
  <initial_audio_gain>0</initial_audio_gain>
  <avsync_enable>true</avsync_enable>
  <user_data></user_data>
  <input_buffer_size>60</input_buffer_size>
  <submitted>2013-01-01 11:06:10 -0800</submitted>
  <status>complete</status>
  <average_fps>26.75</average_fps>
  <buffer_avg>1</buffer_avg>
  <buffer_max>51</buffer_max>
  <dropped_frames>0</dropped_frames>
  <start_time>2013-01-02 18:27:46 -0800</start_time>
  <complete_time>2013-01-03 09:10:05 -0800</complete_time>
  <elapsed>52929</elapsed>
  <audit_messages>
    <audit>
      <code>10</code>
      <created_at>2013-01-02T18:27:50-08:00</created_at>
      <message>Initial timecode: --:--:--:--</message>
    </audit>
    <audit>
      <code>10</code>
      <created_at>2013-01-03T09:10:00-08:00</created_at>
      <message>Final timecode: --:--:--:--</message>
    </audit>
  </audit_messages>
</live_event>
  <next href="http://10.129.232.160/live_events?page=2&amp;per_page=20"/>
</live_event_list>`

func main() {
	http.HandleFunc("/live_events", liveEventList)
	http.HandleFunc("/live_events/", returnBody)

	err := http.ListenAndServe(":8009", nil)
	if err != nil {
		fmt.Println("Failed to start elemental emulator server on :8009")
	}
}

func liveEventList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/xml")
	fmt.Fprint(w, xmlString)
}

func returnBody(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body.")
	}
	defer r.Body.Close()

	fmt.Println("Recieved a cuepoint directed at ", r.URL.String())
	fmt.Println(string(body))

}
