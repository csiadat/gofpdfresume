package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	indent = 61
)

func main() {
	//Create Resume pdf + Add Page 1
	resume := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := resume.GetPageSize()
	fmt.Printf("width = %v, height = %v", w, h)
	resume.AddPage()

	//Header - Name
	resume.SetFont("Arial", "", 16.5)
	resume.SetTextColor(0, 0, 0)
	_, lineHt := resume.GetFontSize()
	resume.Text(w/2.0-56, lineHt*3.5, "Camelia Siadat")

	//Header - Email, Linkedin, Phone Number
	resume.SetFont("Arial", "", 11)
	_, lineHt = resume.GetFontSize()
	resume.MoveTo(w/2.0-76, 61)
	resume.MultiCell(153, lineHt*1.15, "csiadat@scu.edu\nlinkedin.com/in/camelia-siadat\n408-(concealed for privacy)", gofpdf.BorderNone, gofpdf.AlignCenter, false)

	//Dividing Line
	resume.MoveTo(indent, 61+3.75*lineHt*1.15)
	resume.LineTo(550, 61+3.75*lineHt*1.15)
	resume.ClosePath()
	resume.SetFillColor(0, 0, 0)
	resume.SetLineWidth(0.2)
	resume.DrawPath("DF")

	//Objective
	resume.MoveTo(indent, 110+lineHt/2)
	resume.SetFont("Arial", "B", 11)
	_, lineHt = resume.GetFontSize()
	resume.CellFormat(50, lineHt, "Objective", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")
	resume.SetFont("Arial", "", 11)
	resume.MoveTo(indent, 110+lineHt*1.75)
	resume.CellFormat(50, lineHt, "Computer Science Internship", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")

	//Summary
	resume.SetFont("Arial", "B", 11)
	_, lineHt = resume.GetFontSize()
	resume.MoveTo(indent, resume.GetY()+lineHt*1.15)
	resume.CellFormat(50, lineHt, "Summary", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")
	resume.SetFont("Arial", "", 11)
	resume.MoveTo(indent, resume.GetY()+2.5)
	_, lineHt = resume.GetFontSize()
	resume.MultiCell(500, lineHt*1.15, "As a Junior Computer Science student at Santa Clara University (SCU), I'm searching for an internship to apply and enhance my programming knowledge. My experience with various programming languages and hardware systems, combined with my leadership skills, makes me confident in taking on programming challenges during an internship.", gofpdf.BorderNone, gofpdf.AlignLeft, false)

	//Education
	resume.SetFont("Arial", "B", 11)
	_, lineHt = resume.GetFontSize()
	resume.MoveTo(indent, resume.GetY()+lineHt*1.15)
	resume.CellFormat(50, lineHt, "Education", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")
	resume.SetFont("Arial", "", 11)
	resume.MoveTo(indent, resume.GetY()+2.5)
	_, lineHt = resume.GetFontSize()
	resume.MultiCell(500, lineHt*1.15, "Santa Clara University, Santa Clara, CA\nB.S. In Computer Science (Data Science Emphasis) - Dec 2024", gofpdf.BorderNone, gofpdf.AlignLeft, false)

	//Technical Skills
	resume.SetFont("Arial", "B", 11)
	_, lineHt = resume.GetFontSize()
	resume.MoveTo(indent, resume.GetY()+lineHt*1.15)
	resume.CellFormat(50, lineHt, "Technical Skills", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")

	//Technical Skills Bulleted List
	var techSkills = [2]string{"Programming Languages - Scratch, Javascript, Java, Python, C/C++, VerilogHDL, Assembly, Scala.", "Tools - Greenfoot, BlueJ, Atom, Intel FPGA Quartus Prime Design, Questa - Intel FPGA Edition, XCode, Microsoft Visual Studio, Linux, AWS, John the Ripper, Wireshark."}
	y := resume.GetY() + lineHt
	for i := 0; i < len(techSkills); i++ {
		y = bulletedList(resume, "Arial", "", 11, techSkills[i], indent, y)
	}

	//Projects
	resume.SetFont("Arial", "B", 11)
	_, lineHt = resume.GetFontSize()
	resume.MoveTo(indent, resume.GetY()+lineHt*1.15)
	resume.CellFormat(50, lineHt, "Projects", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")

	//Projects Bulleted List
	var projects = [5]string{"Computer Security - Coded and deployed(using C, Python, and AWS) various attacks, including SETUID, Buffer Overflow, and String Format attacks.", "Algorithms - Devised various algorithms by applying techniques including but not limited to Decrease and Conquer, Huffman Encoding, Dijkstra's Algorithm, and Decision Trees.", "Data Structures - Developed a program (using C++, Unix, and Atom text editor) to create a binary tree of pluses, multiplication symbols, and integers; this program would add or multiply the integers to output a final product.", "Introduction to Logic Design - Created a programmable counter, using Verilog, Intel's Quartus Prime Design Software, and an Intel DE2-115 board; this program would continue to count from a predetermined initial value to a maximum value, with pause and resume functionality.", "Introduction to Embedded Systems - Created three Assembly functions that would convert a date into a day of the week using Zeller's Rule, using the following guidelines: using no MUL instructions, using no SDIV/UDIV instructions, and using no MUL instructions and no SDIV/UDIV instructions."}
	y = resume.GetY() + lineHt
	for i := 0; i < len(projects); i++ {
		y = bulletedList(resume, "Arial", "", 11, projects[i], indent, y)
	}

	//Extracurriculars
	resume.SetFont("Arial", "B", 11)
	_, lineHt = resume.GetFontSize()
	resume.MoveTo(indent, resume.GetY()+lineHt*1.15)
	resume.CellFormat(50, lineHt, "Extracurriculars", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")

	//Extracurriculars Bulleted List Pt 1
	var ecs = [2]string{"SCU Creative, Augmented, and Virtual Environments (CAVE) Laboratory - SwingBeats - Create, test, and debug haptic feedback ankle bracelets (HFABs) by soldering HFAB circuitry and running Arduino test programs.", "SCU Ethical, Pragmatic, and Intelligent Computing (EPIC) Laboratory - TailorEd - Assisted in the development of a tool (used by SCU professors) that helps identify classroom environments and teaching styles that best support student learning."}
	y = resume.GetY() + lineHt
	for j := 0; j < len(ecs); j++ {
		y = bulletedList(resume, "Arial", "", 11, ecs[j], indent, y)
	}

	//drawGrid(resume)

	resume.AddPage()
	resume.MoveTo(indent, 61)
	y = 61
	var ecs2 = [3]string{"SCU Society of Women Engineers - SWE++ Committee - Promoted gender inclusivity and diversity by teaching Scratch and Python to middle school girls.", "University of Michigan (via Coursera) - Python for Everybody Specialization.", "SCU Middle Eastern North African Club Board Member - Contributed to the creation of a safe space for MENA-identifying students and allies through hosting informative meetings and events surrounding MENA culture and current events."}
	for j := 0; j < len(ecs2); j++ {
		y = bulletedList(resume, "Arial", "", 11, ecs2[j], indent, y)
	}

	//Professional Experience
	resume.SetFont("Arial", "B", 11)
	_, lineHt = resume.GetFontSize()
	resume.MoveTo(indent, y+lineHt)
	resume.CellFormat(w, lineHt, "Professional Experience", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")
	resume.MoveTo(indent, resume.GetY()+lineHt)
	resume.SetFont("Arial", "", 11)
	_, lineHt = resume.GetFontSize()
	resume.CellFormat(w, lineHt, "Santa Clara University, Community Facilitator", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")
	resume.MoveTo(w-3.1*indent, resume.GetY()-lineHt)
	resume.CellFormat(w, lineHt, "August 2022 - June 2023", gofpdf.BorderNone, 1, gofpdf.AlignLeft, false, 0, "")

	//Professional Experience Bulleted List
	var proexp = [3]string{"Created a safe and comfortable environment for students living in SCU's Graham Residence Hall.", "Planned quarterly events, such as community conversations or self-care nights, to build community and ease new students' transition into university life.", "Member of the Dining Operations Committee - Contributed to the improvement and expansion of dietary options."}
	y = resume.GetY() + 2
	for k := 0; k < len(proexp); k++ {
		y = bulletedList(resume, "Arial", "", 11, proexp[k], indent, y)
	}

	//Open pdf and Check for Output Error
	err := resume.OutputFileAndClose("resume.pdf")
	if err != nil {
		panic(err)
	}
}

// Draw a grid over the resume, to keep track of the positioning of all resume sections
func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("Arial", "", 12)
	pdf.SetLineWidth(1)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}

// My savior aka the function that creates every bulleted list in my resume
func bulletedList(pdf *gofpdf.Fpdf, font string, fontStyle string, fontSize float64, bullet string, curX float64, curY float64) float64 {
	w, _ := pdf.GetPageSize()
	pdf.SetFont(font, fontStyle, fontSize)
	_, lineHt := pdf.GetFontSize()
	pdf.MoveTo(curX, curY)
	pdf.MultiCell(500, lineHt*1.15, "- "+bullet, gofpdf.BorderNone, gofpdf.AlignLeft, false)
	temp := pdf.SplitLines([]byte(bullet), w/1.74)
	maximumY := curY + float64(len(temp))*lineHt - lineHt
	if maximumY > curY {
		curY = maximumY
	}
	curY = curY + lineHt*1.15
	return curY
}
