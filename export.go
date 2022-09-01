package data

import (
	"io"
	"io/fs"
	"os"
	"path"
)

type ExportSpeaker struct {
	Name string
	Img  string
	Desc []string
}

type ExportPartner struct {
	Name string
	Img  string
	Url  string
}

// Expect array of partners data to extract, path where to copy all expected images and prefix for images path
func GetSpeakers(imSp []ImportSpeaker, dest string, prefix string) (expSp []ExportSpeaker, err error) {
	var proceedSpeaker speaker
	var proceedDesc []string
	var outSpeaker ExportSpeaker

	var img string
	var imgToCopy []string

	for _, sp := range imSp {
		proceedSpeaker = speakers[sp.Id]

		proceedDesc = []string{}
		for _, i := range sp.Desc {
			proceedDesc = append(proceedDesc, proceedSpeaker.Desc[i])
		}

		img = proceedSpeaker.Img[sp.Img]
		imgToCopy = append(imgToCopy, img)

		outSpeaker = ExportSpeaker{
			Name: proceedSpeaker.Name[sp.Name],
			Img:  path.Join(prefix, img),
			Desc: proceedDesc,
		}
		expSp = append(expSp, outSpeaker)
	}
	if err = os.MkdirAll(dest, os.ModePerm); err == nil {
		err = copyImg(dest, imgs, "img/speakers", imgToCopy)
	}
	return
}

// Expect array of speakers data to extract, path where to copy all expected images and prefix for images path
func GetPartners(imPa []ImportPartner, dest string, prefix string) (expPa []ExportPartner, err error) {
	var proceedPartner partner
	var outPartner ExportPartner

	var img string
	var imgToCopy []string

	for _, pa := range imPa {
		proceedPartner = partners[pa.Id]

		img = proceedPartner.Img[pa.Img]
		imgToCopy = append(imgToCopy, img)

		outPartner = ExportPartner{
			Name: proceedPartner.Name[pa.Name],
			Img:  path.Join(prefix, img),
			Url:  proceedPartner.Url[pa.Url],
		}
		expPa = append(expPa, outPartner)
	}
	if err = os.MkdirAll(dest, os.ModePerm); err == nil {
		err = copyImg(dest, imgs, "img/partners", imgToCopy)
	}
	return
}

func copyImg(dst string, src fs.FS, prefix string, imgToCopy []string) (err error) {
	for _, i := range imgToCopy {
		if err = copy(i, dst, src, prefix); err != nil {
			break
		}
	}
	return
}

func copy(name, dst string, src fs.FS, prefix string) (err error) {
	var source fs.File
	var destination *os.File
	if source, err = src.Open(path.Join(prefix, name)); err != nil {
		return
	}
	defer source.Close()

	if destination, err = os.Create(path.Join(dst, name)); err != nil {
		return
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return
}
