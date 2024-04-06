## r0sbag CLI

Alternative `rosbag` utility packed with powerful features.

### Installation

Install with:

```
go install
```

### Usage

#### Describe a bag file

```
$ r0sbag info demo.bag
  path:          /home/wyatt/data/bags/demo.bag           
  version:       2.0                                      
  duration:      7.78s                                    
  start:         Mar 21 2017 19:26:20.10 (1490149580.10)  
  end:           Mar 21 2017 19:26:27.88 (1490149587.88)  
  size:          67.1 MB                                  
  messages:      1606                                     
  compression:   lz4 [79/79 chunks; 55.47%]               
  uncompressed:  120.9 MB @ 15.5 MB/s                     
  compressed:    67.1 MB @ 8.6 MB/s (55.47%)              
  types:     diagnostic_msgs/DiagnosticArray  [60810da900de1dd6ddd437c3503511da]       
             sensor_msgs/CompressedImage      [8f7a12909da2c9d3332d540a0977563f]       
             tf2_msgs/TFMessage               [94810edda583a504dfda3829e70d7eec]       
             sensor_msgs/PointCloud2          [1158d486dd51d683ce2f1be655c3c181]       
             sensor_msgs/Range                [c005c34273dc426c67a020a87bc24148]       
             radar_driver/RadarTracks         [6a2de2f790cb8bb0e149d45d297462f8]       
  topics:    /diagnostics              52 msgs    : diagnostic_msgs/DiagnosticArray    
             /image_color/compressed  234 msgs    : sensor_msgs/CompressedImage        
             /tf                      774 msgs    : tf2_msgs/TFMessage                 
             /radar/points            156 msgs    : sensor_msgs/PointCloud2            
             /radar/range             156 msgs    : sensor_msgs/Range                  
             /radar/tracks            156 msgs    : radar_driver/RadarTracks           
             /velodyne_points          78 msgs    : sensor_msgs/PointCloud2
```


#### List the chunks in a bag

```
$ r0sbag list chunks ~/data/bags/demo.bag | head
  offset    start                end                  connections  messages  
  4117      1490149580103843113  1490149580113944947  7            8         
  703705    1490149580124028613  1490149580217237988  7            22        
  1614425   1490149580227348197  1490149580309379447  6            20        
  2525224   1490149580319458697  1490149580410323613  6            20        
  3418190   1490149580420411238  1490149580507495905  6            16        
  4311034   1490149580517578405  1490149580608392239  6            20        
  5213189   1490149580618484655  1490149580711475822  6            23        
  6109581   1490149580721562030  1490149580814003739  6            21        
  7018125   1490149580824118364  1490149580914902989  6            20

```

#### Convert a bag to JSON

```
$ r0sbag cat ~/data/bags/demo.bag | head -n 1
{"topic": "/diagnostics", "time": 1490149580103843113, "data": {"header":{"seq":2602,"stamp":1490149580.113375843,"frame_id":""},"status":[{"level":0,"name":"velodyne_nodelet_manager: velodyne_packets topic status","message":"Desired frequency met; Timestamps are reasonable.","hardware_id":"Velodyne HDL-32E","values":[{"key":"Events in window","value":"100"},{"key":"Events since startup","value":"26020"},{"key":"Duration of window (s)","value":"10.008710"},{"key":"Actual frequency (Hz)","value":"9.991298"},{"key":"Target frequency (Hz)","value":"9.988950"},{"key":"Minimum acceptable frequency (Hz)","value":"8.990055"},{"key":"Maximum acceptable frequency (Hz)","value":"10.987845"},{"key":"Earliest timestamp delay:","value":"0.000300"},{"key":"Latest timestamp delay:","value":"0.000322"},{"key":"Earliest acceptable timestamp delay:","value":"-1.000000"},{"key":"Latest acceptable timestamp delay:","value":"5.000000"},{"key":"Late diagnostic update count:","value":"0"},{"key":"Early diagnostic update count:","value":"0"},{"key":"Zero seen diagnostic update count:","value":"0"}]}]}}
```

#### Print message timestamps

```
$ r0sbag cat ~/data/bags/demo.bag --simple | head -n 1
1490149580103843113 /diagnostics [diagnostic_msgs/DiagnosticArray] [42 10 0 0 204 224 209 88 99 250]...
```

See `r0sbag -h` for available commands.
