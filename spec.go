package framer

type Interface interface {
	// Day returns full day frames for the configured period.
	//
	//     [
	//       {
	//         sta: 2022-05-14 00:00:00
	//         end: 2022-05-15 00:00:00
	//       },
	//       {
	//         sta: 2022-05-15 00:00:00
	//         end: 2022-05-16 00:00:00
	//       },
	//       {
	//         sta: 2022-05-16 00:00:00
	//         end: 2022-05-17 00:00:00
	//       }
	//     ]
	//
	Day() []Frame
	// Exa returns full day frames for the configured period, on the edges
	// exactly bound to the provided start and end time.
	//
	//     [
	//       {
	//         sta: 2022-05-14 15:33:45
	//         end: 2022-05-15 00:00:00
	//       },
	//       {
	//         sta: 2022-05-15 00:00:00
	//         end: 2022-05-16 00:00:00
	//       },
	//       {
	//         sta: 2022-05-16 00:00:00
	//         end: 2022-05-17 01:10:30
	//       }
	//     ]
	//
	Exa() []Frame
	// Lat returns full day frames for the configured period, on the edges
	// exactly bound to the provided end time.
	//
	//     [
	//       {
	//         sta: 2022-05-14 00:00:00
	//         end: 2022-05-15 00:00:00
	//       },
	//       {
	//         sta: 2022-05-15 00:00:00
	//         end: 2022-05-16 00:00:00
	//       },
	//       {
	//         sta: 2022-05-16 00:00:00
	//         end: 2022-05-17 01:10:30
	//       }
	//     ]
	//
	Lat() []Frame
}
